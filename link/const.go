package link

const (
	// Get Method
	Get = "GET"
	// GoH2Pushed is the header key used to transmit Pushed Link header values
	GoH2Pushed = "Go-H2-Pushed"
	// headerAmountLimit is the max number of header values parsed
	headerAmountLimit = 64
	// http2ErrPushLimitReached is the err returned when too many pushes were send : https://golang.org/src/net/http/h2_bundle.go
	http2ErrPushLimitReached = "http2: push would exceed peer's SETTINGS_MAX_CONCURRENT_STREAMS"
	// http2ErrRecursivePush is the err returned when a recursive push was send : https://golang.org/src/net/http/h2_bundle.go
	http2ErrRecursivePush = "http2: recursive push not allowed"
	// Link Header Key
	Link = "Link"
	// XForwardedFor Header Key
	XForwardedFor = "X-Forwarded-For"
)
