package helpers

import "net/url"

// Check the validity of the inputs
func IsInvalidInput(url url.Values, str ...string) string {
	for _, v := range str {
		if url.Get(v) == "" {
			return v
		}
	}
	return ""
}
