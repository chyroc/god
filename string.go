package god

import (
	"strings"
)

// CountPre .
func CountPre(s, substr string) int {
	if s == "" || substr == "" {
		return 0
	}
	if strings.HasPrefix(s, substr) {
		return CountPre(s[len(substr):], substr) + 1
	}
	return 0
}
