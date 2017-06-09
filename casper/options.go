package casper

type settings struct {
	cookieMaxAge                  int
	inferCookieMaxAgeFromResponse bool
}

// Option -> Functional Options
type Option func(*settings)

// DefaultCookieMaxAge sets the cookie maxage
func DefaultCookieMaxAge(maxAge int) func(*settings) {
	return func(s *settings) {
		s.cookieMaxAge = maxAge
	}
}

// InferCookieMaxAgeFromResponse will set cookie maxage = http.ResponseWrite.Header() maxage value (if found)
func InferCookieMaxAgeFromResponse() func(*settings) {
	return func(s *settings) {
		s.inferCookieMaxAgeFromResponse = true
	}
}
