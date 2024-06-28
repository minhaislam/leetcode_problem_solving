package main

import "fmt"

func main() {
	// val := 2
	// nums := []int{1, 1, 2, 3}
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	// val := 0
	nums := []int{3, 4, 3, 3, 1, 6, 1}

	// output := []int{1, 2, 2, 3, 5, 6}

	majorityElement(nums)

}

func majorityElement(nums []int) int {
	// fmt.Println(len(nums) / 2)
	mejority_element := nums[0]

	current_count := 1
	// final_count := 0
	i := 1
	for i < len(nums) {
		if current_count == 0 {
			mejority_element = nums[i]
			current_count = 0
		}

		if mejority_element == nums[i] {
			current_count++
		} else {
			current_count--
		}

		i++
	}
	// j := 0
	// for j < len(nums) {
	// 	if mejority_element == nums[j] {
	// 		final_count++
	// 	}
	// 	j++
	// }

	fmt.Println(mejority_element)
	return mejority_element
}
