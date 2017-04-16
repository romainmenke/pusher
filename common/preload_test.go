package common

import "testing"

func TestCSS(t *testing.T) {
	c := CSS("/stylesheet.css")
	if c.LinkHeader() != "</stylesheet.css>; rel=preload; as=style;" {
		t.Fail()
	}

	if c.Path() != "/stylesheet.css" {
		t.Fail()
	}

	if c.String() != "/stylesheet.css" {
		t.Fail()
	}
}

func TestJS(t *testing.T) {
	j := JS("/bundle.js")
	if j.LinkHeader() != "</bundle.js>; rel=preload; as=script;" {
		t.Fail()
	}

	if j.Path() != "/bundle.js" {
		t.Fail()
	}

	if j.String() != "/bundle.js" {
		t.Fail()
	}
}

func TestFont(t *testing.T) {
	f := Font("/pretty.woff")
	if f.LinkHeader() != "</pretty.woff>; rel=preload; as=font;" {
		t.Fail()
	}

	if f.Path() != "/pretty.woff" {
		t.Fail()
	}

	if f.String() != "/pretty.woff" {
		t.Fail()
	}
}

func TestImg(t *testing.T) {
	i := Img("/trees.jpg")
	if i.LinkHeader() != "</trees.jpg>; rel=preload; as=image;" {
		t.Fail()
	}

	if i.Path() != "/trees.jpg" {
		t.Fail()
	}

	if i.String() != "/trees.jpg" {
		t.Fail()
	}
}
