package main

import "fmt"

func main() {
	num_1 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	nums_2 := []int{1, 2, 2}
	m1 := 6
	n1 := 3
	// output := []int{1, 2, 2, 3, 5, 6}

	merge(num_1, m1, nums_2, n1)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// fmt.Println(nums1, m, nums2, n, len(nums1))
	// k := 0
	k := m
	for k < len(nums1) {

		// fmt.Println(nums2[k-m])

		nums1[k] = nums2[k-m]

		k++

	}
	// fmt.Println(nums1)

	for i := len(nums1) - 1; i >= 0; i-- {
		// fmt.Println(nums1[i])
		for j := 0; j <= i-1; j++ {
			// fmt.Println("2nd Loop", nums1[j])
			if nums1[j] > nums1[j+1] {
				swap_value := nums1[j]
				nums1[j] = nums1[j+1]
				nums1[j+1] = swap_value

			}
		}

	}
	fmt.Println(nums1)
}

// func replace_values(nums1 []int, index_val int) {

// }
