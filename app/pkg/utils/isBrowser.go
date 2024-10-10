package utils

import "strings"

// isBrowser checks if the User-Agent indicates a browser
func IsBrowser(userAgent string) bool {
	// Basic checks for common browsers
	browserAgents := []string{
		"Mozilla", // Common to all browsers
		"Chrome",  // Chrome
		"Safari",  // Safari
		"Firefox", // Firefox
		"Edge",    // Edge
	}

	for _, agent := range browserAgents {
		if strings.Contains(userAgent, agent) {
			return true
		}
	}
	return false
}
