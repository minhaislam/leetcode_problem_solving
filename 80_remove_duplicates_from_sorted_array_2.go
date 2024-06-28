package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 1, 1, 2, 2, 3}

	removeDuplicates2(nums)

}

func removeDuplicates2(nums []int) int {
	index_count := 0

	i := 1
	max_occurance := 1
	previous_index := 1
	// target_index := 0
	for i < len(nums) {
		fmt.Println(nums[i])
		// max_occurance++

		if nums[i] == nums[i-1] {
			max_occurance++
		} else {
			max_occurance = 1
		}

		if max_occurance <= 2 {
			nums[previous_index] = nums[i]
			previous_index++
		}
		i++

	}
	fmt.Println(previous_index)

	return index_count

}
