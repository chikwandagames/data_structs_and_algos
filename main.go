package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	// nums := []int{3, 2, 4}
	var target = 9

	r := twoSum(nums, target)
	fmt.Println(r)

}

func twoSum(nums []int, target int) []int {
	size := len(nums)
	var result []int

	var i, j int

	for i = 0; i < size; i++ {
		for j = 0; j < size; j++ {
			if i != j {
				if nums[i]+nums[j] == target {
					return append(result, i, j)
				}
			}
		}
	}

	return result

}
