package linkheader

import (
	"bufio"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"
)

type HeaderMux struct {
	mu    *sync.RWMutex
	m     map[string]muxEntry
	links map[string]struct{}
	hosts bool // whether any patterns contain hostnames
}

type muxEntry struct {
	explicit bool
	headers  []string
	pattern  string
}

func New() *HeaderMux {
	return &HeaderMux{
		m:     make(map[string]muxEntry),
		links: make(map[string]struct{}),
		mu:    &sync.RWMutex{},
	}
}

func (mux *HeaderMux) Read(filePath string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currentPath := ""
	currentHeaders := []string{}

	for scanner.Scan() {

		txt := scanner.Text()
		if len(txt) > 0 && txt[:1] == "/" {
			currentHeaders = []string{}
			currentPath = txt
			continue
		}
		if currentPath != "" && len(txt) > 0 && txt[:1] == "<" {
			currentHeaders = append(currentHeaders, txt)
			continue
		}
		if txt == "" {
			if currentPath != "" && len(currentHeaders) > 0 {
				mux.Store(currentPath, currentHeaders)
			}
			currentPath = ""
			currentHeaders = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	return nil

}

func (mux *HeaderMux) Store(pattern string, headers []string) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if pattern == "" {
		panic("http: invalid pattern " + pattern)
	}
	if mux.m[pattern].explicit {
		panic("http: multiple registrations for " + pattern)
	}

	for _, h := range headers {
		link := parseLinkHeader(h)
		if link == "" {
			continue
		}
		mux.links[link] = struct{}{}
	}

	if mux.m == nil {
		mux.m = make(map[string]muxEntry)
	}
	mux.m[pattern] = muxEntry{explicit: true, headers: headers, pattern: pattern}

	if pattern[0] != '/' {
		mux.hosts = true
	}

	// Helpful behavior:
	// If pattern is /tree/, insert an implicit permanent redirect for /tree.
	// It can be overridden by an explicit registration.
	n := len(pattern)
	if n > 0 && pattern[n-1] == '/' && !mux.m[pattern[0:n-1]].explicit {
		// If pattern contains a host name, strip it and use remaining
		// path for redirect.
		path := pattern
		if pattern[0] != '/' {
			// In pattern, at least the last character is a '/', so
			// strings.Index can't be -1.
			path = pattern[strings.Index(pattern, "/"):]
		}
		url := &url.URL{Path: path}
		mux.m[pattern[0:n-1]] = muxEntry{headers: headers, pattern: url.String()}
	}
}

// Does path match pattern?
func pathMatch(pattern, path string) bool {
	if len(pattern) == 0 {
		// should not happen
		return false
	}
	n := len(pattern)
	if pattern[n-1] != '/' {
		return pattern == path
	}
	return len(path) >= n && path[0:n] == pattern
}

func cleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	// path.Clean removes trailing slash except for root;
	// put the trailing slash back if necessary.
	if p[len(p)-1] == '/' && np != "/" {
		np += "/"
	}
	return np
}

func (mux *HeaderMux) match(path string) (headers []string, pattern string) {

	_, found := mux.links[path]
	if found {
		return []string{}, path
	}

	var n = 0
	for k, v := range mux.m {
		if !pathMatch(k, path) {
			continue
		}
		if headers == nil || len(headers) == 0 || len(k) > n {
			n = len(k)
			headers = v.headers
			pattern = v.pattern
		}
	}
	return
}

func (mux *HeaderMux) SetLinkHeaders(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		return
	}

	headers, _ := mux.getHeaders(r.Host, r.URL.Path)
	for _, h := range headers {
		w.Header().Add("Link", h)
	}

}

func (mux *HeaderMux) getHeaders(host, path string) (headers []string, pattern string) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()

	// Host-specific pattern takes precedence over generic ones
	if mux.hosts {
		headers, pattern = mux.match(host + path)
	}
	if headers == nil {
		headers, pattern = mux.match(path)
	}
	if headers == nil {
		headers, pattern = []string{}, ""
	}
	return
}

func parseLinkHeader(h string) string {

	var linkStart int
	var linkEnd int

RUNELOOP:
	for index, runeValue := range h {
		switch runeValue {
		case '<':
			linkStart = index + 1
		case '>':
			linkEnd = index
			break RUNELOOP
		case ';':
			linkStart = 0
			linkEnd = 0
		}
	}

	if linkStart == 0 || linkEnd == 0 {
		return ""
	}

	if !strings.Contains(h[linkEnd+1:], "rel=preload") {
		return ""
	}
	if strings.Contains(h[linkEnd+1:], "nopush") {
		return ""
	}

	res := strings.TrimSpace(h[linkStart:linkEnd])

	if isAbsolute(res) {
		return ""
	}

	return res
}

func isAbsolute(p string) bool {
	if strings.Contains(p, "http://") || strings.Contains(p, "https://") || strings.Contains(p, "www.") {
		return true
	}

	// if a '.' is found before '/' it is an absolute url
	// requesting "stylesheet.css" instead of "/stylesheet.css" is not ok and not supported
	for _, runeValue := range p {
		switch runeValue {
		case '.':
			return true
		case '/':
			return false
		}
	}

	return false
}
