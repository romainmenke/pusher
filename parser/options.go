package parser

import "time"

type settings struct {
	cacheTime time.Duration
	withCache bool
}

// Option -> Functional Options
type Option func(*settings)

// WithCache makes the Handler cache which assets were pushed/preloaded after parsing the html response once.
// Useful html files that don't change after server start
func WithCache() func(*settings) {
	return func(s *settings) {
		s.withCache = true
	}
}

// CacheDuration sets the cache duration
func CacheDuration(d time.Duration) func(*settings) {
	return func(s *settings) {
		s.cacheTime = d
	}
}
