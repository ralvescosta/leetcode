package main

import (
	"strings"
)

func LengthOfLastWord(s string) int {
	if len(s) == 1 && s != " " {
		return 1
	}

	truncated := strings.Split(s, " ")

	for i := len(truncated) - 1; i >= 0; i-- {
		t := strings.Trim(truncated[i], "")
		if t != "" {
			return len(truncated[i])
		}
	}

	return -1
}
