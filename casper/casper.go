// Copyright (c) 2017 Taichi Nakashima
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package casper

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/romainmenke/pusher/casper/internal/encoding/golomb"
	"github.com/romainmenke/pusher/casper/internal/intsort"
)

const (
	// defaultCookieName is default cookie name for storing
	// a fingerprint of asset files being cached by the browser.
	defaultCookieName = "X-Go-Casper"

	// defaultCookiePath is default cookie path to be used for
	// generating cookie to return.
	defaultCookiePath = "/"
)

// Casper provides a interface for cache-aware HTTP/2 server push.
type Casper struct {
	p uint
	n uint

	// skipPush decides executing actual server push or not. This should
	// be used only in testing.
	//
	// Currently, it's kinda hard to receive http push in go http client.
	// This should be removed in future.
	skipPush bool
}

type hasher struct {
	hash.Hash
	b []byte
}

func (h *hasher) Close() {
	h.Reset()
	h.b = h.b[:8]
	hasherPool.Put(h)
}

func getHasher() *hasher {
	h := hasherPool.Get().(*hasher)
	h.Reset()
	h.b = h.b[:8]
	return h
}

// hasherPool is the sync.Pool used to reduce GC pauses
var hasherPool *sync.Pool

// init initialises the writerPool
func init() {
	hasherPool = &sync.Pool{
		New: func() interface{} {
			return &hasher{
				md5.New(),
				make([]byte, 8, 8),
			}
		},
	}
}

// hash generate a hash value from the given bytes for
// n elements and p faslse positive probability.
//
// It's ok to use md5 since we just need a hash that generates
// uniformally-distributed values for best results.
func (c *Casper) hash(p []byte) uint {
	h := getHasher()
	defer h.Close()

	h.Write(p)
	b := h.Sum(nil)
	hex.Encode(h.b, b[12:16])
	i, err := strconv.ParseUint(string(h.b), 16, 32)
	if err != nil {
		panic(err)
	}

	return uint(i) % (c.n * c.p)
}

type encoder struct {
	io.WriteCloser
	buf *bytes.Buffer
}

func (e *encoder) Close() {
	e.buf.Reset()
	encoderPool.Put(e)
}

func getEncoder() *encoder {
	e := encoderPool.Get().(*encoder)
	e.buf.Reset()

	return e
}

// encoderPool is the sync.Pool used to reduce GC pauses
var encoderPool *sync.Pool

// init initialises the writerPool
func init() {
	encoderPool = &sync.Pool{
		New: func() interface{} {
			buf := bytes.Buffer{}
			return &encoder{
				WriteCloser: base64.NewEncoder(base64.RawURLEncoding, &buf),
				buf:         &buf,
			}
		},
	}
}

// generateCookie generates cookie from the given hash values.
func (c *Casper) generateCookie(hashValues []uint) (*http.Cookie, error) {

	// golomb encoder expect the given array is sorted.
	sort.Sort(intsort.Uints(hashValues))

	encoder := getEncoder()
	defer encoder.Close()

	if err := golomb.Encode(encoder, hashValues, c.p); err != nil {
		return nil, fmt.Errorf("failed golomb coding: %s", err)
	}

	if err := encoder.WriteCloser.Close(); err != nil {
		return nil, fmt.Errorf("failed to close encoder: %s", err)
	}

	return &http.Cookie{
		Name:   defaultCookieName,
		Value:  encoder.buf.String(),
		MaxAge: 3600,
		Path:   defaultCookiePath,
	}, nil
}

// readCookie reads cookie from http request and decode it to hash array.
func (c *Casper) readCookie(r *http.Request, hashValues *[]uint) error {
	cookie, err := r.Cookie(defaultCookieName)
	if err != nil && err != http.ErrNoCookie {
		return fmt.Errorf("failed to read cookie: %s", err)
	}

	if err == http.ErrNoCookie {
		return nil
	}

	// Decode golomb coded cookie value to original hash values array.
	decoder := base64.NewDecoder(base64.RawURLEncoding, strings.NewReader(cookie.Value))
	err = golomb.DecodeAll(decoder, c.p, hashValues)
	if err != nil {
		return fmt.Errorf("failed golomb decoding: %s", err)
	}

	return nil
}

// search looks up the provided slices contains the given value.
func search(a []uint, h uint) bool {
	for i := 0; i < len(a); i++ {
		if h == a[i] {
			return true
		}

		if h < a[i] {
			return false
		}
	}
	return false
}
