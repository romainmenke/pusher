package link

const (
	GoH2Pushed               = "Go-H2-Pushed"
	XForwardedFor            = "X-Forwarded-For"
	Link                     = "Link"
	Get                      = "GET"
	headerAmountLimit        = 50
	headerLengthLimit        = 1500
	http2ErrRecursivePush    = "http2: recursive push not allowed"
	http2ErrPushLimitReached = "http2: push would exceed peer's SETTINGS_MAX_CONCURRENT_STREAMS"
)
