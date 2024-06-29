package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// Code to measure

	nums := []int{3}

	// output := []int{1, 2, 2, 3, 5, 6}

	canJump(nums)
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())

}

func canJump(nums []int) bool {
	i := len(nums) - 2
	target_index := len(nums) - 1
	status := true
	// target_index := len(nums) - 1
	// if len
	// dead_end := 0
	// dead_end_index := 0
	for i >= 0 {
		if nums[i]+i >= target_index && nums[i] > 0 {
			fmt.Println(target_index, "---", nums[i], " index ", i)
			target_index = i
			status = true
			fmt.Println(status)
		} else {
			fmt.Println(target_index, "---", nums[i], " index ", i)
			status = false
			fmt.Println(status)

		}
		i--
	}
	fmt.Println(status)
	return status
}
