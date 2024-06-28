package main

import "fmt"

func main() {
	// val := 2
	// nums := []int{1, 1, 2, 3}
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	// val := 0
	nums := []int{3}

	// output := []int{1, 2, 2, 3, 5, 6}

	majorityElement(nums)

}

func majorityElement(nums []int) int {
	fmt.Println(len(nums) / 2)
	current_count := 0
	final_element := 0
	if len(nums) == 1 {
		current_count = 0
		final_element = nums[0]
	}
	i := 0
	// fmt.Println("Going to count occurance for value: ", nums[i])
	for len(nums) > 1 && i < len(nums) {
		start_val := nums[i]
		current_count = 1
		j := i + 1
		fmt.Println("Going to count occurance for value: ", nums[i])
		for j < len(nums) {
			if nums[j] == start_val {
				current_count++
			}
			j++
		}
		fmt.Println("total Occurance: ", current_count)

		if current_count > len(nums)/2 {
			final_element = start_val
			break
		} else {
			current_count = 0
			final_element = start_val
		}
		i++

	}
	fmt.Println(current_count)
	return final_element

}
