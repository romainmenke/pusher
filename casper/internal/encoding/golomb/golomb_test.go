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
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strconv"
	"testing"

	"github.com/romainmenke/pusher/casper/internal/bits"
	"github.com/romainmenke/pusher/casper/internal/intsort"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		input []byte
		p     uint
		want  uint
		err   error
	}{
		{
			[]byte{0xcb, 0x80}, // 11001011 10000000
			1 << 6,
			151,
			nil,
		},
		{
			[]byte{0xcb}, // 11001011
			1 << 5,
			75,
			io.EOF,
		},
		{
			[]byte{0x00}, // 00000000
			1 << 7,
			0,
			errPadding,
		},
	}

	for _, tc := range cases {
		rd := bytes.NewReader(tc.input)
		br := bits.NewReader(rd)
		got, err := decode(br, tc.p)
		if err != tc.err {
			t.Errorf("error=%v, want=%v", err, tc.err)
		}

		if got != tc.want {
			t.Errorf("decode=%v, want=%v", got, tc.want)
		}

	}
}

func TestDecodeAll(t *testing.T) {
	cases := []struct {
		input []byte
		p     uint
		want  []uint
	}{
		{
			[]byte{0xcb, 0x80}, // 11001011 10000000
			1 << 6,
			[]uint{151},
		},

		{
			[]byte{0xcb, 0xcf}, // 11001011 11001111
			1 << 5,
			[]uint{75, 154},
		},

		{
			[]byte{0x81, 0x4e},
			1 << 6,
			[]uint{65, 104},
		},
	}

	for _, tc := range cases {
		rd := bytes.NewReader(tc.input)
		got, err := DecodeAll(rd, tc.p)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Decode=%v, want=%v", got, tc.want)
		}

	}
}

func TestEncoding(t *testing.T) {
	cases := []struct {
		input []uint
		p     uint
		want  []byte
	}{
		{
			[]uint{151},
			1 << 6,
			[]byte{0xcb, 0x80}, // 11001011 10000000
		},

		{
			[]uint{65, 104},
			1 << 6,
			[]byte{0x81, 0x4e}, // 10000001 1001110
		},
	}

	for _, tc := range cases {
		var buf bytes.Buffer
		if err := Encode(&buf, tc.input, tc.p); err != nil {
			t.Fatal(err)
		}

		if got := buf.Bytes(); !bytes.Equal(got, tc.want) {
			t.Errorf("Encode=%x, want=%x", got, tc.want)
		}
	}
}

func ExampleEncode() {
	// Number of elements and false positive probability.
	//
	// Minimum number of bits is N*log(P)
	// = 26 * log(1<<6) = 156 bits = 19 bytes
	N, P := uint(26), uint(1<<6)

	// Example data set comes from https://github.com/rasky/gcs
	file := "./testdata/words.nato"
	f, _ := os.Open(file)
	sc := bufio.NewScanner(f)
	defer f.Close()

	hashF := func(v []byte) uint {
		h := md5.New()
		h.Write(v)
		b := h.Sum(nil)

		s := hex.EncodeToString(b[12:16])
		i, err := strconv.ParseUint(s, 16, 32)
		if err != nil {
			panic(err)
		}
		return uint(i) % (N * P)
	}

	a := make([]uint, 0, N)
	for sc.Scan() {
		a = append(a, hashF(sc.Bytes()))
	}
	sort.Sort(intsort.Uints(a))

	// Encode hash value array to Golomb-coded sets
	// and write it to buffer.
	var buf bytes.Buffer
	Encode(&buf, a, P)

	fmt.Printf("%x", buf.Bytes())
	// Output: cba920f780663a061f2065198ab1032d624c50331e66ae9818
}
