package parser

import (
	"net/http"

	"github.com/romainmenke/pusher/common"
	"golang.org/x/net/html"
)

func (w *responseWriter) extractLinks() chan common.Preloadable {

	out := make(chan common.Preloadable)

	go func() {
		defer close(out)

		links := make(map[common.Preloadable]struct{})
		preloads := make(map[string]struct{})

		contentType := http.DetectContentType(w.body.Bytes())
		if contentType != "text/html; charset=utf-8" {
			return
		}

		path := w.request.URL.RequestURI()

		z := html.NewTokenizer(w.body)

	TOKENIZER:
		for {
			tt := z.Next()

			var asset common.Preloadable
			var preload string

			switch tt {
			case html.ErrorToken:
				// End of the document, we're done
				break TOKENIZER
			case html.SelfClosingTagToken:

				t := z.Token()
				asset, preload = parseToken(t, path)

				if asset != nil {
					if _, found := preloads[asset.Path()]; !found {
						links[asset] = struct{}{}
						asset = nil
					}
				} else if preload != "" {
					preloads[preload] = struct{}{}
					preload = ""
				}

			case html.StartTagToken:

				t := z.Token()
				asset, preload = parseToken(t, path)

				if asset != nil {
					if _, found := preloads[asset.Path()]; !found {
						links[asset] = struct{}{}
						asset = nil
					}
				} else if preload != "" {
					preloads[preload] = struct{}{}
					preload = ""
				}

			}
		}

		index := 0
		for key := range links {
			if _, found := preloads[key.Path()]; found {
				continue
			}
			if index >= common.HeaderAmountLimit {
				return
			}

			if w.settings.withCache {
				putOneInCache(path, key)
			}
			out <- key
		}
	}()

	return out
}

const (
	hrefStr   = "href"
	imgStr    = "img"
	linkStr   = "link"
	relStr    = "rel"
	scriptStr = "script"
	srcStr    = "src"
)

func parseToken(t html.Token, path string) (common.Preloadable, string) {

	var (
		asset     common.Preloadable
		isPreload bool
	)

	switch t.Data {
	case linkStr:

		for _, attr := range t.Attr {
			switch attr.Key {
			case relStr:
				if attr.Val == common.Preload {
					isPreload = true
				}
			case common.NoPush:
				return nil, ""
			case hrefStr:
				if common.IsAbsolute(attr.Val) || attr.Val == path {
					return nil, ""
				}
				asset = common.CSS(attr.Val)
			}
		}

	case scriptStr:

		for _, attr := range t.Attr {
			switch attr.Key {
			case relStr:
				if attr.Val == common.Preload {
					return nil, ""
				}
			case common.NoPush:
				return nil, ""
			case srcStr:
				if common.IsAbsolute(attr.Val) || attr.Val == path {
					return nil, ""
				}
				asset = common.JS(attr.Val)
			}
		}

	case imgStr:

		for _, attr := range t.Attr {
			switch attr.Key {
			case relStr:
				if attr.Val == common.Preload {
					return nil, ""
				}
			case common.NoPush:
				return nil, ""
			case srcStr:
				if common.IsAbsolute(attr.Val) || attr.Val == path {
					return nil, ""
				}
				asset = common.Img(attr.Val)
			}
		}
	}

	if isPreload {
		return nil, asset.Path()
	}

	return asset, ""
}
