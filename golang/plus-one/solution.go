package main

func PlusOne(digits []int) []int {
	lastDigit := digits[len(digits)-1]

	if lastDigit < 9 {
		digits[len(digits)-1] = lastDigit + 1
		return digits
	}

	lessThanNineIndex := -1

	for i := len(digits) - 2; i >= 0; i-- {
		if digits[i] < 9 {
			lessThanNineIndex = i
			break
		}
	}

	if lessThanNineIndex == -1 {
		res := make([]int, len(digits)+1)
		res[0] = 1
		return res
	}

	digits[lessThanNineIndex] = digits[lessThanNineIndex] + 1
	for i := lessThanNineIndex + 1; i < len(digits); i++ {
		digits[i] = 0
	}

	return digits
}
