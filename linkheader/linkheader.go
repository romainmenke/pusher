package linkheader

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/romainmenke/pusher/common"
)

type settings struct {
	path string
}

// Option -> Functional Options : Awesome
type Option func(*settings)

// RulesFileOption is used to pass the location of your Rules File to the Handler func
func RulesFileOption(path string) func(*settings) {
	return func(s *settings) {
		s.path = path
	}
}

// Handler wraps an http.Handler.
// It sets Link Headers for paths declared in the Rules File.
func Handler(handler http.Handler, options ...Option) http.Handler {

	s := &settings{}
	m := &sync.RWMutex{}

	for _, option := range options {
		option(s)
	}

	if s.path == "" {
		return handler
	}

	linkMap, assetMap, err := read(s.path)
	if err != nil {
		return handler
	}

	mux := http.NewServeMux()

	for path := range linkMap {
		scopedPath := path
		h := wrap(scopedPath, assetMap, linkMap, m, handler)
		mux.Handle(scopedPath, h)
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
		if txt == "-" {
			continue
		}
		if currentPath != "" && len(txt) > 0 && txt[:1] == "<" {
			link := common.ParseLinkHeader(txt)
			if link == "" {
				continue
			}
			currentHeaders = append(currentHeaders, txt)
			headerMap[link] = struct{}{}
			continue
		}
		if txt == "" {
			if currentPath != "" {
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
