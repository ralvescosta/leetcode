package main

import "strings"

func StrStrEasy(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func StrStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	for i, f := range haystack {
		if f == rune(needle[0]) {
			for j, s := range needle {
				if i+j >= len(haystack) {
					return -1
				}

				if rune(haystack[i+j]) != s {
					break
				}

				if j == len(needle)-1 {
					return i
				}
			}
		}
	}

	return -1
}
