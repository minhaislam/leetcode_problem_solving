package main

import "fmt"

func main() {
	// val := 2
	// nums := []int{0, 1, 2, 2, 3, 0, 4, 2}

	val := 0
	nums := []int{}

	// output := []int{1, 2, 2, 3, 5, 6}

	removeElement(nums, val)

}

// func removeElement(nums []int, val int) int {
// 	// fmt.Println(nums)
// 	new_num := make([]int, len(nums))

// 	i := 0
// 	k := 0
// 	len_v := len(new_num) - 1
// 	for i < len(nums) {
// 		if nums[i] == val {
// 			new_num[len_v] = nums[i]
// 			// fmt.Println(new_num)
// 			len_v--
// 		} else {
// 			new_num[k] = nums[i]
// 			// fmt.Println(new_num)
// 			k++
// 		}
// 		i++
// 	}
// 	nums = new_num
// 	fmt.Println(nums)
// 	return k
// 	// fmt.Println(nums)

// }

func removeElement(nums []int, val int) int {
	// new_num := make([]int, len(nums))

	i := 0
	k := 0
	// len_v := len(new_num) - 1
	for len(nums) > 1 && i < len(nums) {
		if nums[i] != val {
			nums[k] = nums[i]
			// fmt.Println(new_num)
			k++
		}
		i++
	}
	// // nums = new_num
	// fmt.Println(new_num)
	// j := 0
	// for len(nums) > 0 && j < k+1 {
	// 	nums[j] = new_num[j]
	// 	j++
	// }
	fmt.Println(nums)
	// // fmt.Println(new_num)
	return k
}
