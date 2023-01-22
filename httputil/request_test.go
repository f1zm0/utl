package httputil_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/f1zm0/utl/httputil"
	"github.com/stretchr/testify/assert"
)

func TestAddCustomReqHeaders(t *testing.T) {
	testCases := map[string]struct {
		additionalHeaders string
		expected          map[string]string
	}{
		"standard headers": {
			additionalHeaders: "X-Test-Header: test-value, X-Test-Header-2: test-value-2",
			expected: map[string]string{
				"X-Test-Header":   "test-value",
				"X-Test-Header-2": "test-value-2",
			},
		},
		"no headers": {
			additionalHeaders: "",
			expected:          map[string]string{},
		},
		"colon only": {
			additionalHeaders: ":",
			expected:          map[string]string{},
		},
		"new headers": {
			additionalHeaders: "X-Test-Header: test-value",
			expected:          map[string]string{"X-Test-Header": "test-value"},
		},
		"empty headers": {
			additionalHeaders: "X-Test-Header: test-value, , X-Test-Header-2: test-value-2",
			expected: map[string]string{
				"X-Test-Header":   "test-value",
				"X-Test-Header-2": "test-value-2",
			},
		},
		"empty header name": {
			additionalHeaders: ": test-value",
			expected:          map[string]string{},
		},
		"empty header value": {
			additionalHeaders: "X-Test-Header:",
			expected:          map[string]string{"X-Test-Header": ""},
		},
		"headers with spaces": {
			additionalHeaders: " X-Test-Header: test-value   , X-Test-Header-2: test-value-2 ",
			expected: map[string]string{
				"X-Test-Header":   "test-value",
				"X-Test-Header-2": "test-value-2",
			},
		},
		"invalid header format": {
			additionalHeaders: "X-Test-Header: test-value, X-Test-Header-2: test-value-2, X-Test-Header-3",
			expected: map[string]string{
				"X-Test-Header":   "test-value",
				"X-Test-Header-2": "test-value-2",
			},
		},
		"CRLF injection": {
			additionalHeaders: "X-Test-Header: test-value\r\nX-Injected-Header: injected-value",
			expected:          map[string]string{"X-Test-Header": "test-value"},
		},
		"CRLF injection urlencode": {
			additionalHeaders: "X-Test-Header: test-value%0d%0aX-Injected-Header: injected-value",
			expected:          map[string]string{"X-Test-Header": "test-value"},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			r := &http.Request{Header: http.Header{}}
			actualHeaders := map[string]string{}

			httputil.AddCustomReqHeaders(r, tc.additionalHeaders)
			for hn, hv := range r.Header {
				actualHeaders[hn] = strings.Join(hv, ",")
			}

			assert.Equal(t, tc.expected, actualHeaders)
		})
	}
}
