package httputil_test

import (
	"testing"

	"github.com/f1zm0/utl/httputil"
	"github.com/stretchr/testify/assert"
)

func TestIsSameDomain(t *testing.T) {
	type args struct {
		baseURL   string
		targetURL string
	}

	testCases := map[string]struct {
		args     args
		expected bool
	}{
		"same domain": {args: args{"http://a.b", "http://a.b/path/to/test"}, expected: true},
		"diff scheme": {args: args{"https://a.b", "http://a.b/path/to/test"}, expected: true},
		"subdomain-1": {args: args{"http://a.b", "http://x.a.b"}, expected: false},
		"subdomain-2": {args: args{"http://a.b", "http://a.b.c"}, expected: false},
		"diff domain": {args: args{"http://a.b", "http://x.y.z"}, expected: false},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			assert.Equal(t, tc.expected, httputil.IsSameDomain(tc.args.baseURL, tc.args.targetURL))
		})
	}
}
