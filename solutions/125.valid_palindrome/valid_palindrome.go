package validpalindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	sanitized := sanitizeString(s)

	if sanitized != "" && len(sanitized) > 1 {
		fwIdx := 0
		backIdx := len(sanitized) - 1

		for {
			if sanitized[fwIdx] != sanitized[backIdx] {
				return false
			}
			fwIdx++
			backIdx--
			if fwIdx == backIdx || fwIdx > backIdx {
				break
			}
		}
	}

	return true
}

func sanitizeString(s string) string {
	sanitized := strings.ToLower(s)
	return regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(sanitized, "")
}
