package main

func MySqrt(x int) int {
	if x <= 0 {
		return 0
	}

	// Tolerance for the accuracy of the result
	const tolerance = 1e-10
	z := float64(x)
	var prev float64

	for {
		// Update the guess according to Newton's method
		prev = z
		z -= (z*z - float64(x)) / (2 * z)

		diff := z - prev

		// Check if the change is within the desired tolerance
		if diff < 0 {
			diff = -diff
		}

		if diff < tolerance {
			break
		}
	}

	return int(z)
}
