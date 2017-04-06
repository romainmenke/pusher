package linkheader

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
)

type settings struct {
	path string
}

type Option func(*settings)

func PathOption(path string) func(*settings) {
	return func(s *settings) {
		s.path = path
	}
}

func Handler(handler http.Handler, options ...Option) http.Handler {

	s := &settings{}

	for _, option := range options {
		option(s)
	}

	if s.path == "" {
		return handler
	}

	pathMap, headerMap, err := read(s.path)
	if err != nil {
		return handler
	}

	mux := http.NewServeMux()

	for path := range pathMap {
		localPath := path
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "GET" {
				handler.ServeHTTP(w, r)
				return
			}

			if _, found := headerMap[r.URL.Path]; found {
				handler.ServeHTTP(w, r)
				return
			}

			for _, header := range pathMap[localPath] {
				w.Header().Add("Link", header)
			}

			handler.ServeHTTP(w, r)
		})

		mux.Handle(path, h)
	}

	return mux
}

func read(filePath string) (map[string][]string, map[string]struct{}, error) {

	pathMap := make(map[string][]string)
	headerMap := make(map[string]struct{})

	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
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
			headerMap[parseLinkHeader(txt)] = struct{}{}
			continue
		}
		if txt == "" {
			if currentPath != "" && len(currentHeaders) > 0 {
				pathMap[currentPath] = currentHeaders
			}
			currentPath = ""
			currentHeaders = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		if err == io.EOF {
			return pathMap, headerMap, nil
		}
		return nil, nil, err
	}

	return pathMap, headerMap, nil

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

	res := strings.TrimSpace(h[linkStart:linkEnd])

	return res
}
