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

package bits

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	input := []byte{0xff, 0x0f} // 1111 1111 0000 1111
	rd := bytes.NewReader(input)
	reader := NewReader(rd)

	cases := []struct {
		n    int
		want uint
	}{
		{2, 3},  // 11
		{3, 7},  // 111
		{5, 28}, // 11100
		{3, 1},  // 001
		{3, 7},  // 111
	}

	for _, tc := range cases {
		got, err := reader.Read(tc.n)
		if err != nil && err != io.EOF {
			t.Fatalf("Read(%d) should not fail: %s", tc.n, err)
		}

		if got != tc.want {
			t.Errorf("Read(%d)=%b, want=%b", tc.n, got, tc.want)
		}

	}
}

func TestWriter(t *testing.T) {
	cases := []struct {
		size   int
		inputs []uint
		want   []byte
	}{
		{8, []uint{255}, []byte{0xff}},
		{4, []uint{15, 15}, []byte{0xff}},
		{2, []uint{3, 3, 3, 3}, []byte{0xff}},
		{1, []uint{1, 1, 1, 1, 1, 1, 1, 1}, []byte{0xff}},

		{4, []uint{15, 15, 15}, []byte{0xff, 0xf0}},
		{2, []uint{3, 3, 3, 3, 3, 3}, []byte{0xff, 0xf0}},
	}

	for _, tc := range cases {
		var buf bytes.Buffer
		writer := NewWriter(&buf)

		for _, input := range tc.inputs {
			if err := writer.Write(input, tc.size); err != nil {
				t.Fatalf("Write should not fail: %s", err)
			}
		}

		if err := writer.Flush(); err != nil {
			t.Fatalf("Flush should not fail: %s", err)
		}

		if !bytes.Equal(buf.Bytes(), tc.want) {
			t.Errorf("Write writes %x, want %x", buf.Bytes(), tc.want)
		}
	}
}

func TestMask(t *testing.T) {
	cases := []struct {
		input int
		want  string
	}{
		{8, "11111111"},
		{4, "00001111"},
		{2, "00000011"},
	}

	for _, tc := range cases {
		m := mask(tc.input)
		if got := fmt.Sprintf("%08b", m); got != tc.want {
			t.Errorf("mask(%d)=%s,want=%s", tc.input, got, tc.want)
		}
	}
}
