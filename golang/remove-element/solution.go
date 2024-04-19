package main

func RemoveElement(nums []int, val int) int {
	correct := []int{}

	for _, v := range nums {
		if v != val {
			correct = append(correct, v)
		}
	}

	copy(nums, correct)

	return len(correct)
}
