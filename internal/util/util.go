package util

// FallbackString return fallback if value is empty string.
func FallbackString(val, fallback string) string {
	if val == "" {
		return fallback
	}

	return val
}
