package common

import "fmt"

// Preloadable represents a preloadable asset url
type Preloadable interface {
	// LinkHeader returns a usable Link Header
	LinkHeader() string
	// Path returns the relative path to the asset
	Path() string
	// Conform to fmt.Stringer
	String() string
}

// CSS represents a css url
type CSS string

// LinkHeader returns a usable Link Header
func (s CSS) LinkHeader() string {
	return fmt.Sprintf("<%s>; rel=preload; as=style;", string(s))
}

// Path returns the relative path to the asset
func (s CSS) Path() string {
	return string(s)
}

func (s CSS) String() string {
	return s.Path()
}

// JS represents a javascript url
type JS string

// LinkHeader returns a usable Link Header
func (s JS) LinkHeader() string {
	return fmt.Sprintf("<%s>; rel=preload; as=script;", string(s))
}

// Path returns the relative path to the asset
func (s JS) Path() string {
	return string(s)
}

func (s JS) String() string {
	return s.Path()
}

// Font represents a font url
type Font string

// LinkHeader returns a usable Link Header
func (s Font) LinkHeader() string {
	return fmt.Sprintf("<%s>; rel=preload; as=font;", string(s))
}

// Path returns the relative path to the asset
func (s Font) Path() string {
	return string(s)
}

func (s Font) String() string {
	return s.Path()
}

// Img represents an image url
type Img string

// LinkHeader returns a usable Link Header
func (s Img) LinkHeader() string {
	return fmt.Sprintf("<%s>; rel=preload; as=image;", string(s))
}

// Path returns the relative path to the asset
func (s Img) Path() string {
	return string(s)
}

func (s Img) String() string {
	return s.Path()
}
