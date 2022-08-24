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

// FindLastAfterSubstr 寻找字符串 s 中，匹配最后一个 substr 后的字符串
//
// 当 substr 为空的时候，必定返回空字符串
// 当 s==1/2/3 substr=/ 返回 3
func FindLastAfterSubstr(s, substr string) string {
	if substr == "" {
		return ""
	}
	start := len(s) - len(substr)
	end := len(s)
	for start >= 0 {
		if s[start:end] == substr {
			return s[end:]
		}
		start--
		end--
	}

	return ""
}
