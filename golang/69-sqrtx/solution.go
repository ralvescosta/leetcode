package main

func MySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2

	for left <= right {

		mid := (left + right) / 2

		s := mid * mid
		if s == x {
			return mid
		}

		if s > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}
