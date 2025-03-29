package main

func main() {
	println(binary_serach([]int{1, 2, 3, 4, 5}, 20))
}

func binary_serach(vec []int, value int) int {
	p := len(vec) / 2
	for range vec {
		if vec[p] == value {
			return p
		}

		if vec[p] > value {
			p = p / 2
		} else {
			p += p / 2
		}

		if p > len(vec) || p < 0 {
			return -1
		}
	}

	return -1

}
