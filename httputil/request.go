package httputil

import (
	"net/http"
	"strings"
)

// AddCustomReqHeaders parses a comma separated string of
// HTTP headers and adds them to the provided http.Request object.
func AddCustomReqHeaders(req *http.Request, headers string) {
	for _, rawKeyVal := range strings.Split(headers, ",") {

		kv := strings.Split(strings.Trim(rawKeyVal, " "), ":")
		if kv[0] != "" {
			kk := strings.Trim(kv[0], " ")
			vv := strings.Trim(kv[1], " ")
			req.Header.Add(kk, vv)
		}
	}
}
