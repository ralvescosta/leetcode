package main

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, v := range nums {
		comp := target - v

		if idx, founded := hash[comp]; founded {
			return []int{i, idx}
		}

		hash[v] = i
	}

	return []int{-1, -1}
}
