package helper

import "strings"

func GetLastUriSegment(uri string) string {
	parts := strings.Split(uri, "/")
	size := len(parts)
	if size == 1 {
		return parts[0]
	}
	return parts[size-1]
}
