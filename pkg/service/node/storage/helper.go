package storage

import (
	"fmt"
	"strings"
)

func IsValidContent(content string) bool {
	switch content {
	case "iso", "vztmpl":
		return true
	default:
		return false
	}
}

func contentPath(node, storage string) string {
	return fmt.Sprintf("/nodes/%s/storage/%s/content", node, storage)
}

func ParseContentName(s string) string {
	// storage-name:content-name,size=...
	return strings.Split(strings.Split(s, ":")[1], ",")[0]
}
