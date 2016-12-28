package pusher

import (
	"net/http"
	"strings"
)

// isPageGet is used to differentiate between Pushable and Non-Pushable requests.
// It returns true when it is Non-Pushable.
// It uses the file extension in the request URL and the http Method.
func isPageGet(r *http.Request) bool {
	if r.Method != "GET" {
		return false
	}

	p := r.URL.Path
	p = strings.ToLower(p)
	components := strings.Split(p, ".")
	extension := components[len(components)-1]
	_, found := assetExtension[extension]
	if found {
		return false
	}

	return true
}

// assetExtension is a map with file extensions corresponding to files that might be pushable.
var assetExtension = map[string]struct{}{
	// json
	"json": struct{}{},
	// css
	"css": struct{}{},
	// js
	"js": struct{}{},
	// fonts
	"eot":   struct{}{},
	"woff":  struct{}{},
	"woff2": struct{}{},
	"ttf":   struct{}{},
	// image
	"jpeg": struct{}{},
	"jpg":  struct{}{},
	"png":  struct{}{},
	"gif":  struct{}{},
	"tiff": struct{}{},
	"svg":  struct{}{},
	// audio
	"3gp":   struct{}{},
	"aa":    struct{}{},
	"aac":   struct{}{},
	"aax":   struct{}{},
	"act":   struct{}{},
	"aiff":  struct{}{},
	"amr":   struct{}{},
	"ape":   struct{}{},
	"au":    struct{}{},
	"awb":   struct{}{},
	"dct":   struct{}{},
	"dss":   struct{}{},
	"dvf":   struct{}{},
	"flac":  struct{}{},
	"gsm":   struct{}{},
	"iklax": struct{}{},
	"ivs":   struct{}{},
	"m4a":   struct{}{},
	"m4b":   struct{}{},
	"m4p":   struct{}{},
	"mmf":   struct{}{},
	"mp3":   struct{}{},
	"mpc":   struct{}{},
	"msv":   struct{}{},
	"ogg,":  struct{}{},
	"opus":  struct{}{},
	"ra,":   struct{}{},
	"raw":   struct{}{},
	"sln":   struct{}{},
	"tta":   struct{}{},
	"vox":   struct{}{},
	"wav":   struct{}{},
	"wma":   struct{}{},
	"wv":    struct{}{},
	"webm":  struct{}{},
	// movie
	"mkv":  struct{}{},
	"flv":  struct{}{},
	"vob":  struct{}{},
	"ogv":  struct{}{},
	"drc":  struct{}{},
	"gifv": struct{}{},
	"mng":  struct{}{},
	"avi":  struct{}{},
	"mov":  struct{}{},
	"wmv":  struct{}{},
	"yuv":  struct{}{},
	"rm":   struct{}{},
	"rmvb": struct{}{},
	"asf":  struct{}{},
	"amv":  struct{}{},
	"mp4":  struct{}{},
	"m4v":  struct{}{},
	"mpg":  struct{}{},
	"mpeg": struct{}{},
	"mpv":  struct{}{},
	"m2v":  struct{}{},
	"svi":  struct{}{},
	"3g2":  struct{}{},
	"mxf":  struct{}{},
	"roq":  struct{}{},
	"nsv":  struct{}{},
	"f4p":  struct{}{},
	"f4a":  struct{}{},
	"f4b":  struct{}{},
}
