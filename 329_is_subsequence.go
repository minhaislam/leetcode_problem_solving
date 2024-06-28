package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// Code to measure

	s := "aaaaaa"
	t := "bbaaaaa"

	// output := []int{1, 2, 2, 3, 5, 6}

	isSubsequence(s, t)
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())

}

func isSubsequence(s string, t string) bool {

	i := 0
	tag_val := 0
	status := true
	main_string := len(t)
	subsequent_string := len(s)

	if main_string < subsequent_string {
		status = false
	}

	for main_string >= subsequent_string && i < subsequent_string {
		fmt.Println(string(s[i]))
		j := tag_val
		for j < main_string {

			if s[i] == t[j] {
				fmt.Println("the string ", string(s[i]), " found in ", j)
				tag_val = j + 1
				status = true
				break
			} else if s[i] != t[j] {
				fmt.Println("the string ", string(s[i]), " not found in ", j)
				status = false
			}
			j++
		}

		if j >= main_string {
			status = false
			break
		}

		if status == false {
			break
		}
		i++
	}
	fmt.Println(status)
	return status
}
