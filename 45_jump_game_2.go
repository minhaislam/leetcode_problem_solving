// Runtime: 47 ms
// Memory Usage: 7.1 MB

// Runtime beats 66.10% submissions

// Memory beats 81% submissions

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// Code to measure

	nums := []int{2, 3, 0, 1, 4}

	// output := []int{1, 2, 2, 3, 5, 6}

	jump(nums)
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())

}

func jump(nums []int) int {
	i := 0
	// initial_value := nums[0]
	// // step := initial_value
	current_index := 0
	count := 0
	max_val_index := 0
	for i < len(nums)-1 {
		fmt.Println("value: ", nums[i])
		if max_val_index < i+nums[i] {
			max_val_index = i + nums[i]
		}

		if max_val_index >= len(nums)-1 {
			count++
			break
		}
		if i == current_index {
			count++
			current_index = max_val_index
		}

		i++
	}
	// fmt.Println("current_index: ", current_index)
	fmt.Println(count)
	return count
}
