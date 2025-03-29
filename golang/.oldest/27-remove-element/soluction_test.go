package main

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	input := []int{3, 2, 2, 3}
	val := 3

	RemoveElement(input, val)

	fmt.Println(input)
}
