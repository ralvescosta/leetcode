package main

func RemoveDuplicates(nums []int) int {
	lastValue := nums[0]
	i := 1

	for {
		if i >= len(nums) {
			break
		}

		if nums[i] == lastValue {
			nums = append(nums[:i-1], nums[i:]...)
		} else {
			lastValue = nums[i]
			i++
		}
	}

	return len(nums)
}
