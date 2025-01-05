package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	isPalindrome(121)
	fmt.Printf("r: %v", isPalindromeWithoutStrconv(121))
}

func isPalindromeWithoutStrconv(x int) bool {
	//edge cases
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	half := 0
	for x > half {
		half = (half * 10) + (x % 10)
		x = x >> 10
		x = int(math.Floor(float64(x) / 10))
	}

	return x == half || x == int(math.Floor(float64(half)/10))
}

func isPalindrome(x int) bool {
	//edge cases
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	str := strconv.Itoa(x)
	l := len(str)
	desc := l - 1

	for asc := 0; asc < l; asc++ {
		if str[asc] != str[desc] {
			return false
		}
		desc -= 1
	}

	return true
}
