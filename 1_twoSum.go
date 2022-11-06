package main

import "fmt"

func main() {
	var nums = []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
}

func twoSum(nums []int, target int) []int {
	valueIndex := map[int]int{}

	for index, number := range nums {
		if index2, ok := valueIndex[target-number]; ok {
			return []int{index2, index}
		}

		valueIndex[number] = index
	}
	return []int{}
}
