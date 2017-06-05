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

package golomb

import (
	"errors"
	"io"
	"math"

	"github.com/romainmenke/pusher/casper/internal/bits"

	"fmt"
)

var errPadding = errors.New("padding")

// DecodeAll decodes...
func DecodeAll(rd io.Reader, p uint, dst *[]uint) error {
	br := bits.NewReader(rd)
	prev := uint(0)
	for {
		v, err := decode(br, p)
		if err == errPadding {
			// Ignore padding value
			return nil
		}

		if err == io.EOF {
			*dst = append(*dst, v+prev)
			return nil
		}

		if err != nil {
			return err
		}

		*dst = append(*dst, v+prev)

		prev = v + prev
	}
}

func decode(br *bits.Reader, p uint) (uint, error) {
	var v uint

	// Decode unary parts. Sum p until enconter 0 bits.
	for {
		b, err := br.Read(1)
		if err == io.EOF {
			if v == 0 {
				return 0, errPadding
			}
			return 0, fmt.Errorf("unexpected bit format")
		}

		if err != nil {
			return 0, err
		}

		if b == 0 {
			break
		}
		v += p
	}

	// Decode remainder parts.
	bitLen := int(math.Log2(float64(p)))
	r, err := br.Read(bitLen)
	if err == io.EOF && r == 0 {
		return 0, errPadding
	}

	if err != nil && err != io.EOF {
		return 0, err
	}
	v += r

	return v, err
}

// Encode encodes the given uint array and writes to underlying writer.
// p is false-positive probability. The src array must be uniformly
// distribute set of values.
func Encode(w io.Writer, src []uint, p uint) error {
	if len(src) == 0 {
		return nil
	}

	// bitLen is number of bits for writing remainder.
	bitLen := int(math.Log2(float64(p)))

	// TODO(tcnksm)
	wr := bits.NewWriter(w)

	prev := uint(0)
	for _, h := range src {
		v := h - prev
		q, r := v/p, v%p

		// Write unary code of quotient
		if err := wr.Write(1<<(uint(q)+1)-2, int(q+1)); err != nil {
			return err
		}

		// Write remainder
		if err := wr.Write(r, bitLen); err != nil {
			return err
		}

		prev = h
	}

	return wr.Flush()
}
