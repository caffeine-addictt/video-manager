package utils

import "strings"

func Multiline(line ...string) string {
	return strings.Join(line, "\n")
}
