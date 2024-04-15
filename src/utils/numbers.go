package utils

import (
	"fmt"
	"math"
	"strings"
)

// String to int
var StrToInt = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

// Positive Integer Flag
type PositiveIntFlag int

func (i *PositiveIntFlag) String() string {
	return fmt.Sprint(*i)
}

func (i *PositiveIntFlag) Int() int {
	return int(*i)
}

func (i *PositiveIntFlag) Set(value string) error {
	splitV := strings.Split(value, "")
	if len(splitV) == 0 {
		return fmt.Errorf("not a number")
	}

	if splitV[0] == "-" {
		return fmt.Errorf("only positive integers allowed")
	}

	if splitV[0] == "+" {
		splitV = splitV[1:]
	}

	n := 0
	digitR := 0
	for index := len(splitV); index > 0; index-- {
		integer, ok := StrToInt[splitV[index-1]]
		if !ok {
			return fmt.Errorf("not a number")
		}

		n += integer * int(math.Pow10(digitR))
		digitR++
	}

	*i = PositiveIntFlag(n)
	return nil
}

func (i *PositiveIntFlag) Type() string {
	return "<int>"
}
