package rules_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/fd/httpmiddlewarevet"
	"github.com/romainmenke/pusher/rules"
)

func TestMiddlewareWithoutOption(t *testing.T) {
	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return rules.Handler(h)
	})
}

func TestMiddlewareWithFileOption(t *testing.T) {
	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return rules.Handler(h, rules.FileOption("./rules/example/rules.txt"))
	})
}

func TestMiddlewareWithBadFileOption(t *testing.T) {
	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return rules.Handler(h, rules.FileOption("./example/rules.txt"))
	})
}

func TestMiddlewareWithTextOption(t *testing.T) {

	reader := strings.NewReader(`/
</css/stylesheet.css>; rel=preload; as=style;
</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;
/broken_a>; rel=preload;
</broken_b>

/alpha.html
</css/stylesheet.css>; rel=preload; as=style;
</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;
</js/text_change.js>; rel=preload; as=script;

/beta.html
</css/stylesheet.css>; rel=preload; as=style;
</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;
</img/gopher.png>; rel=preload; as=image;
</img/gopher1.png>; rel=preload; as=image;
</img/gopher2.png>; rel=preload; as=image;
</img/gopher3.png>; rel=preload; as=image;
</img/gopher4.png>; rel=preload; as=image;
</img/gopher5.png>; rel=preload; as=image;

/gamma.html
-

/gamma-b.html
</css/stylesheet.css>; rel=preload; as=style;
</fonts/CutiveMono-Regular.ttf>; rel=preload; as=font;
</call.json>; rel=preload;

`)

	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return rules.Handler(h, rules.ReaderOption(reader))
	})
}
