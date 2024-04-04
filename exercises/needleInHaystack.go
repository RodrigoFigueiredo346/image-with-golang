package exercises

import (
	"strings"
)

func StrStr(haystack string, needle string) int {

	x := len(haystack)

	switch {
	case x == 0:
		return -1
	case x < len(needle):
		return -1
	case x == 1:
		if x == len(needle) && haystack == needle {
			return 0
		}

	}

	firstIndex := strings.Index(haystack, needle)

	return firstIndex

}
