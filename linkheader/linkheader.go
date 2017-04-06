package linkheader

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Handler(handler http.Handler, path string) http.Handler {

	options, _ := Read(path)
	mux := http.NewServeMux()

	for path := range options {
		localPath := path
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(localPath, r.URL.String())

			for _, header := range options[localPath] {
				w.Header().Add("Link", header)
			}

			handler.ServeHTTP(w, r)
		})

		mux.Handle(path, h)
	}

	return mux
}

func Read(filePath string) (map[string][]string, error) {

	out := make(map[string][]string)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
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
				out[currentPath] = currentHeaders
			}
			currentPath = ""
			currentHeaders = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		if err == io.EOF {
			return out, nil
		}
		return nil, err
	}

	return out, nil

}
