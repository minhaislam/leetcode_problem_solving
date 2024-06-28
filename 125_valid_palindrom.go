package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	// Code to measure

	s := "A man, a plan, a canal: Panama"

	// output := []int{1, 2, 2, 3, 5, 6}

	isPalindrome(s)
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())

}

func isPalindrome(s string) bool {
	set_bool := true
	i := 0
	k := len(s) - 1
	for i < k {
		if !isAlphanumeric(s[i]) {
			i++
			continue
		}
		if !isAlphanumeric(s[k]) {
			k--
			continue
		}
		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[k])) {
			set_bool = false
			break
		}
		i++
		k--
	}
	fmt.Println(set_bool)
	return set_bool
}

func isAlphanumeric(char_val byte) bool {
	if char_val >= 'A' && char_val <= 'Z' {
		return true
	} else if char_val >= 'a' && char_val <= 'z' {
		return true
	} else if char_val >= '0' && char_val <= '9' {
		return true
	}
	return false
}
