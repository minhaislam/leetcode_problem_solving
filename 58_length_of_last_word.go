package main

func main() {

	str_val := "luffy is still joyboy"

	// output := []int{1, 2, 2, 3, 5, 6}

	lengthOfLastWord(str_val)

}

func lengthOfLastWord(s string) int {
	// fmt.Println(s)
	i := len(s) - 1
	new_str := 0
	vote := -1
	// vote1 := 0
	for i >= 0 {
		// fmt.Println(string(s[i]))
		if string(s[i]) != " " {
			new_str = new_str + 1
			vote = 1
			// vote2 = 0

		} else if vote == 1 && string(s[i]) == " " {
			break
		}

		i--
	}
	// fmt.Println(new_str)
	return new_str
}
