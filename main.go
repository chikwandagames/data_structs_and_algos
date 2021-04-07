package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	// nums := []int{3, 2, 4}
	var target = 9

	r := twoSum(nums, target)
	fmt.Println(r)

}

func twoSum(nums []int, target int) []int {

	var result []int

	for i := 0; i < len(nums); i++ {
		var seen []int
		// Loop from the end len(nums) -1
		for j := len(nums) - 1; j > -1; j-- {
			// fmt.Printf("i: %v \n", i)
			// fmt.Printf("j: %v \n", j)
			if i != j {
				seen = append(seen, j)
			}

		}

		fmt.Println(seen)
		for _, val := range seen {
			// fmt.Println("i:", nums[i])
			// fmt.Println("val: ", nums[val])
			if nums[i]+nums[val] == target {
				// fmt.Println("yesss")
				return append(result, i, val)
			}
		}
	}

	return result
}
