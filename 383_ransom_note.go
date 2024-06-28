package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// Code to measure

	ransomNote := "aa"
	magazine := "aab"

	// output := []int{1, 2, 2, 3, 5, 6}

	canConstruct(ransomNote, magazine)
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())

}

func canConstruct(ransomNote string, magazine string) bool {
	ransom_length := len(ransomNote)
	magazine_length := len(magazine)
	status := true
	m := make(map[string]int)
	i := 0
	for i < magazine_length {
		m[string(magazine[i])] = m[string(magazine[i])] + 1
		i++
	}
	fmt.Println(m)

	j := 0
	for j < ransom_length {
		if m[string(ransomNote[j])] > 0 {
			m[string(ransomNote[j])]--
			status = true
		} else if m[string(ransomNote[j])] == 0 {
			status = false
			break
		}
		j++
	}
	fmt.Println(status)
	return status
}
