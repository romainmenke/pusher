package link

const (
	// Get Method
	Get = "GET"
	// GoH2Pushed is the header key used to transmit Pushed Link header values
	GoH2Pushed               = "Go-H2-Pushed"
	headerAmountLimit        = 64
	headerLengthLimit        = 2048
	http2ErrPushLimitReached = "http2: push would exceed peer's SETTINGS_MAX_CONCURRENT_STREAMS"
	http2ErrRecursivePush    = "http2: recursive push not allowed"
	// Link Header Key
	Link = "Link"
	// NoPush directive
	NoPush = "nopush"
	// Preload directive
	Preload = "rel=preload"
	// XForwardedFor Header Key
	XForwardedFor = "X-Forwarded-For"
)
