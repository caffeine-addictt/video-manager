package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Multiline(line ...string) string {
	return strings.Join(line, "\n")
}

func InputPrompt(text string) string {
	// Tailing space need not be added
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, text+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
