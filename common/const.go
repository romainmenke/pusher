package common

const (
	// H2Pushed is the header key used to transmit Pushed Link header values
	H2Pushed = "H2-Pushed"
	// HeaderAmountLimit is the max number of header values parsed
	HeaderAmountLimit = 64
	// PushAmountLimit is the max number of pushes sent
	PushAmountLimit = 64
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
	// RelIsPreload directive quoted
	RelIsPreloadQuoted = `rel="preload"`
	// XForwardedFor Header Key
	XForwardedFor = "X-Forwarded-For"
)
