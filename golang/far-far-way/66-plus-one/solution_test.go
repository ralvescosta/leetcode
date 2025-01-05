package main

import "testing"

func TestPlusOne(t *testing.T) {
	expected := []int{2, 4, 9, 4, 0}

	res := PlusOne([]int{2, 4, 9, 3, 9})

	for i, v := range expected {
		if res[i] != v {
			t.Errorf("expected %d, got %d", v, res[i])
		}
	}
}
