package httputil

import "net/url"

// IsSameDomain returns a boolean that indicates if the provided URLs have
// the same domain, according to the rules that define the same origin policy (RFC 6454).
func IsSameDomain(baseURL, targetURL string) bool {
	bURL, err := url.Parse(baseURL)
	if err != nil {
		return false
	}

	tURL, err := url.Parse(targetURL)
	if err != nil {
		return false
	}

	if bURL.Host == "" || tURL.Host == "" {
		return false
	}

	return bURL.Host == tURL.Host
}
