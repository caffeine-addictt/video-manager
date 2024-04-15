package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// Strategy Enum
type StrategyEnum string

const (
	StrategySynchronous StrategyEnum = "synchronous"
	StrategyConcurrent  StrategyEnum = "concurrent"
)

func (e *StrategyEnum) String() string {
	return string(*e)
}

func (e *StrategyEnum) Set(value string) error {
	switch value {
	case "concurrent", "synchronous":
		*e = StrategyEnum(value)
		return nil
	default:
		return errors.New("must be one of 'synchronous' or 'concurrent'")
	}
}

func (e *StrategyEnum) Type() string {
	return "<concurrent|synchronous>"
}

// Functions
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

func SimilarityScore(s1, s2 string) int {
	// Calculate how similar two strings are [0-100]
	if s1 == s2 {
		return 100
	}

	similar := 0
	least := min(len(s1), len(s2))
	for i := 0; i < least; i++ {
		if s1[i] == s2[i] {
			similar++
		}
	}

	if least == 0 {
		least = 1
	}

	return similar * 100 / least
}
