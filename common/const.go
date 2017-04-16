package common

const (
	// Get Method
	Get = "GET"
	// GoH2Pushed is the header key used to transmit Pushed Link header values
	GoH2Pushed = "Go-H2-Pushed"
	// HeaderAmountLimit is the max number of header values parsed
	HeaderAmountLimit = 64
	// HeaderLengthLimit is the max length of a header value, longer values are skipped
	HeaderLengthLimit = 1024
	// Link Header Key
	Link = "Link"
	// NoPush directive
	NoPush = "nopush"
	// Preload directive
	Preload = "preload"
	// RelIsPreload directive
	RelIsPreload = "rel=preload"
	// XForwardedFor Header Key
	XForwardedFor = "X-Forwarded-For"
)
