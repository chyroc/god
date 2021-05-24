package god

import (
	"strings"
)

var Str = &str{}

type str struct {
}

func (r *str) HasPrefixList(s string, prefixList ...string) bool {
	for _, prefix := range prefixList {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}
