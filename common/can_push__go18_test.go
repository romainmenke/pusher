// +build go1.8

package common_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/romainmenke/pusher/common"
)

func TestCanPush(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header:     http.Header{},
	}
	var writer http.ResponseWriter
	writer = &testWriter{
		httptest.NewRecorder(),
	}

	if !common.CanPush(writer, request) {
		t.Fail()
	}
}

func TestCanPush_H1(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 1,
		Header:     http.Header{},
	}
	var writer http.ResponseWriter
	writer = &testWriter{
		httptest.NewRecorder(),
	}

	if common.CanPush(writer, request) {
		t.Fail()
	}
}

func TestCanPush_Forwarded(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header: http.Header{
			"X-Forwarded-For": []string{"foo"},
		},
	}
	var writer http.ResponseWriter
	writer = &testWriter{
		httptest.NewRecorder(),
	}

	if common.CanPush(writer, request) {
		t.Fail()
	}
}

func TestCanPush_NoPusher(t *testing.T) {
	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header:     http.Header{},
	}
	var writer http.ResponseWriter
	writer = httptest.NewRecorder()

	if common.CanPush(writer, request) {
		t.Fail()
	}
}

func TestCanPush_NoGet(t *testing.T) {
	request := &http.Request{
		Method:     "POST",
		ProtoMajor: 2,
		Header:     http.Header{},
	}
	var writer http.ResponseWriter
	writer = &testWriter{
		httptest.NewRecorder(),
	}

	if common.CanPush(writer, request) {
		t.Fail()
	}
}

func BenchmarkCanPush(b *testing.B) {

	request := &http.Request{
		Method:     "GET",
		ProtoMajor: 2,
		Header:     http.Header{},
	}

	var writer http.ResponseWriter

	writer = &testWriter{
		httptest.NewRecorder(),
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if common.CanPush(writer, request) {
			continue
		}
	}
}

type testWriter struct {
	http.ResponseWriter
}

func (w *testWriter) Push(target string, opts *http.PushOptions) error {
	return nil
}
