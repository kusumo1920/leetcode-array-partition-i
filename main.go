package main

import (
	"slices"
)

func main() {
	input := []int{1, 4, 3, 2}
	output := arrayPairSum(input)
	println(output)
}

func arrayPairSum(nums []int) int {
	slices.Sort(nums)
	result := 0
	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}
	return result
}
