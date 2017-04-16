package link

const (
	// http2ErrPushLimitReached is the err returned when too many pushes were send : https://golang.org/src/net/http/h2_bundle.go
	http2ErrPushLimitReached = "http2: push would exceed peer's SETTINGS_MAX_CONCURRENT_STREAMS"
	// http2ErrRecursivePush is the err returned when a recursive push was send : https://golang.org/src/net/http/h2_bundle.go
	http2ErrRecursivePush = "http2: recursive push not allowed"
)
