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

	nums := []int{2, 3, 1, 1, 4}

	// output := []int{1, 2, 2, 3, 5, 6}

	jump(nums)
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())

}

func jump(nums []int) int {
	i := len(nums) - 2
	target_index := len(nums) - 1
	count := 0
	// target_index := len(nums) - 1
	// if len
	// dead_end := 0
	// dead_end_index := 0
	for i >= 0 {
		if nums[i]+i >= target_index && nums[i] > 0 {
			fmt.Println(target_index, "---", nums[i], " index ", i)
			target_index = i
			count++
			fmt.Println(count)
		} else {
			fmt.Println(target_index, "---", nums[i], " index ", i)
			count--
			fmt.Println(count)

		}
		i--
	}
	fmt.Println(count)
	return count
}
