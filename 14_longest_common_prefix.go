package main

import "fmt"

func main() {

	strs := []string{"dog", "racecar", "car"}

	// output := []int{1, 2, 2, 3, 5, 6}

	longestCommonPrefix(strs)

}

func longestCommonPrefix(strs []string) string {
	i := 0
	str_val := ""
	// for i < len(strs[0]) {

	// 	str_val = str_val + string(strs[0][i])
	// 	fmt.Println(str_val)
	// 	i++
	// }

	for i < len(strs[0]) {

		str_val = str_val + string(strs[0][i])
		upvote := 1
		// fmt.Println(str_val)
		j := 1
		for j < len(strs) {
			// fmt.Println("comparing ", string(strs[j][:i+1]), " with ", str_val)
			// fmt.Println("the string is  ", len(string(strs[j][:i+1])), " ge than ", len(strs[j]))
			if len(str_val) > len(strs[j]) || string(strs[j][:i+1]) != str_val {
				// fmt.Println("the string is  ", len(str_val), " ge than 11 ", len(strs[j]))
				upvote = 0
				str_val = string(strs[0][:i])
				break
			}
			j++
		}
		if upvote == 0 {
			break
		}

		i++
	}
	fmt.Println(str_val)

	return str_val
}
