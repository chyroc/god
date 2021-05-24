package god_string

import (
	"strings"
)

// HasPrefixList .
func HasPrefixList(s string, prefixList ...string) bool {
	for _, prefix := range prefixList {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

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

// FindLastSubstr .
func FindLastSubstr(s, substr string) string {
	if substr == "" {
		return ""
	}
	v := strings.Split(s, substr)
	return v[len(v)-1] // assert: len must >= 1
}
