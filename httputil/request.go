package httputil

import (
	"net/http"
	"strings"
)

// AddCustomReqHeaders parses a comma separated string of
// HTTP headers and adds them to the provided http.Request object.
func AddCustomReqHeaders(req *http.Request, headers string) {
	for _, rawKeyVal := range strings.Split(headers, ",") {
		var (
			finalKey   string
			finalValue string
		)

		// if there's no : in the header, it's invalid
		if !strings.Contains(rawKeyVal, ":") {
			continue
		}

		// remove spaces and other invalid chars
		keyValStr := strings.TrimSpace(rawKeyVal)
		keyValStr = strings.Split(keyValStr, "\r")[0]
		keyValStr = strings.Split(keyValStr, "\n")[0]
		keyValStr = strings.Split(keyValStr, "%")[0]

		// split header name and value
		kv := strings.Split(keyValStr, ":")

		switch len(kv) {
		case 0:
			// invalid header format
			continue
		case 1:
			// if header name is also empty, skip
			if kv[0] == "" {
				continue
			}
			// if header value is empty, set header name as key and empty string as value
			finalKey = strings.TrimSpace(kv[0])
			finalValue = ""
		default: // len(kv) >= 2
			// if header name is empty, skip
			if kv[0] == "" {
				continue
			}
			// if header value is empty, set header name as key and empty string as value
			finalKey = strings.TrimSpace(kv[0])
			finalValue = strings.TrimSpace(kv[1])
		}

		req.Header.Add(finalKey, finalValue)
	}
}
