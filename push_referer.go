package pusher

import "net/http"

const PushRefererKey = "push-referer"

func isPushedResource(request *http.Request) bool {
	return request.Header.Get(PushRefererKey) != "" // if the push referer header is set, this is pushed content
}

func setInitiatorForOptions(request *http.Request, opts *http.PushOptions) *http.PushOptions {

	if opts.Header == nil {
		opts.Header = http.Header{}
	}

	pushReferer := request.Header.Get(PushRefererKey)

	if pushReferer == "" {
		opts.Header.Add(PushRefererKey, resourceKeyFromRequest(request))
		return opts
	}

	opts.Header.Add(PushRefererKey, pushReferer)
	return opts
}

func setInitiatorForWriter(writer http.ResponseWriter, request *http.Request) string {

	pushReferer := request.Header.Get(PushRefererKey)

	if pushReferer == "" {
		return ""
	}

	writer.Header().Add(PushRefererKey, pushReferer)
	return pushReferer
}
