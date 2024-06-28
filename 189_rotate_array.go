package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53}
	// k := 82

	rotate(nums, k)

}

func rotate(nums []int, k int) {
	fmt.Println("Testcase: ", nums)
	fmt.Println("K:", k)
	array_len := len(nums)
	// if k > len(nums) {
	k = k % array_len
	fmt.Println(k)
	i := 0
	for i < array_len/2 {
		fmt.Println(nums[i])
		tmp_val := nums[i]
		nums[i] = nums[array_len-i-1]
		nums[array_len-i-1] = tmp_val
		i++
	}
	fmt.Println(nums)
	i = 0
	for i < k/2 {
		fmt.Println("2nd round -->", nums)
		tmp_val := nums[i]
		nums[i] = nums[k-i-1]
		nums[k-i-1] = tmp_val
		i++
	}
	fmt.Println(nums)

	i = k
	fmt.Println("diff: ", (array_len - k))
	val := (array_len - k) / 2
	allocate := 0
	fmt.Println("val:", val)
	for i <= array_len-val-1 {
		fmt.Println(nums[i])
		tmp_val := nums[i]
		nums[i] = nums[array_len-allocate-1]
		nums[array_len-allocate-1] = tmp_val
		allocate++
		i++
		fmt.Println("after each rotation", nums)
	}

	fmt.Println(nums)

}
