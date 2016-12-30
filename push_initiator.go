package pusher

import (
	"net/http"
	"strings"
)

const PushInitiatorKey = "push-initiator"

func getInitiator(request *http.Request) string {

	var pushInitiators []string

	for k, v := range request.Header {
		if strings.ToLower(k) == PushInitiatorKey {
			pushInitiators = v
		}
	}

	if len(pushInitiators) == 0 {
		return pathFromReferer(request.Referer())
	}

	return pathFromReferer(pushInitiators[0])
}

func setInitiatorForOptions(request *http.Request, opts *http.PushOptions) *http.PushOptions {

	var pushInitiators []string

	for k, v := range request.Header {
		if strings.ToLower(k) == PushInitiatorKey {
			pushInitiators = v
		}
	}

	if len(pushInitiators) == 0 {
		if opts.Header == nil {
			opts.Header = http.Header{}
		}
		opts.Header.Add(PushInitiatorKey, request.RequestURI)
		return opts
	}

	if opts.Header == nil {
		opts.Header = http.Header{}
	}
	opts.Header.Add(PushInitiatorKey, pushInitiators[0])
	return opts
}

func setInitiatorForWriter(writer http.ResponseWriter, request *http.Request) string {

	var pushInitiators []string

	for k, v := range request.Header {
		if strings.ToLower(k) == PushInitiatorKey {
			pushInitiators = v
		}
	}

	if len(pushInitiators) == 0 {
		return ""
	}

	writer.Header().Add(PushInitiatorKey, pushInitiators[0])
	return pushInitiators[0]
}
