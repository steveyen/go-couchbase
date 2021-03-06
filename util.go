package couchbase

import "strings"

// Return the hostname with the given suffix removed.
func CleanupHost(h, commonSuffix string) string {
	if strings.HasSuffix(h, commonSuffix) {
		return h[:len(h)-len(commonSuffix)]
	}
	return h
}

// Find the longest common suffix from the given strings.
func FindCommonSuffix(input []string) string {
	rv := ""
	if len(input) < 2 {
		return ""
	}
	from := input
	for i := len(input[0]); i > 0; i-- {
		common := true
		suffix := input[0][i:]
		for _, s := range from {
			if !strings.HasSuffix(s, suffix) {
				common = false
				break
			}
		}
		if common {
			rv = suffix
		}
	}
	return rv
}
