package main

func main() {
	// val := 2
	// nums := []int{1, 1, 2, 3}
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	// val := 0
	nums := []int{1, 1, 2}

	// output := []int{1, 2, 2, 3, 5, 6}

	removeDuplicates(nums)

}

func removeDuplicates(nums []int) int {
	i := 0
	reindex := 0
	dup_index := 0
	if len(nums) == 1 {
		reindex++
	}
	for len(nums) > 1 && i < len(nums) {
		// fmt.Println("Value is ", nums[i], " and Index ", i)

		j := 0
		count := 0
		// fmt.Println("finding duplicates for", nums[i])
		for i > 0 && j < i {

			// fmt.Println("Inside 2nd loop")
			// fmt.Println(nums[j], " is ge to ", nums[i])
			if nums[j] == nums[i] {
				// fmt.Println("found Duplicate in index", j)
				// fmt.Println(nums[i], " is Dupliacate & its index", i)
				// fmt.Println("breaking inner loop")
				// fmt.Println("Duplicate found")
				dup_index = j + 1
				// fmt.Println("allocated index: ", dup_index)
				// fmt.Println("Before Count increment:", count)
				count++
				break
			}
			j++
		}
		// fmt.Println("Value of count: ", count)
		if count == 0 {
			// fmt.Println("index of duplicate: ", dup_index)
			tmp := nums[i]
			nums[i] = nums[dup_index]
			nums[dup_index] = tmp
			dup_index = dup_index + 1
			reindex++
		}
		// fmt.Println(i)
		i++
		// fmt.Println("Value after each loop:", nums)
		// fmt.Println("duplicate index value:", dup_index)

		// fmt.Println("-----------")
	}
	// fmt.Println(nums)
	// fmt.Println(reindex)
	return reindex

}
