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

	nums := []int{11, 15}

	// output := []int{1, 2, 2, 3, 5, 6}

	hIndex(nums)
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())

}

func hIndex(citations []int) int {
	// h_index := 0
	// current_index_val := 0
	i := 0
	step := 0
	fmt.Println(len(citations))

	for i < len(citations) {
		j := i + 1
		for j < len(citations) {
			if citations[i] < citations[j] {
				tmp_val := citations[i]
				citations[i] = citations[j]
				citations[j] = tmp_val

			}
			j++
		}

		i++
	}

	k := 0
	for k < len(citations) {
		if citations[k] < k+1 {
			break
		} else {
			step++
		}
		k++
	}

	fmt.Println(citations)
	fmt.Println("Step: ", step)
	return step
}
