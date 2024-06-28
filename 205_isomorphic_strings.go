package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// Code to measure

	s := "aaeaa"
	t := "uuxyy"

	// output := []int{1, 2, 2, 3, 5, 6}

	isIsomorphic(s, t)
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())

}

func isIsomorphic(s string, t string) bool {
	string_length := len(s)
	i := 0
	status := true
	for i < string_length {
		j := i + 1
		// previous_val := t[i]
		// tag := 1
		for j < string_length {
			// fmt.Println(i, " Index Runnig for :", string(s[i]), " previous:", string(previous_val), " and other string ", string(t[j]))
			if (s[i] != s[j] && t[j] == t[i]) || (s[i] == s[j] && t[j] != t[i]) {
				status = false
				break
			}
			j++
		}
		if status == false {
			break
		}
		i++
	}

	// fmt.Println(m)
	// fmt.Println(status)
	return status
}

// func isIsomorphic(s string, t string) bool {
// 	m := make(map[string]string)

// 	i := 0
// 	status := true
// 	for i < len(s) {
// 		if val, ok := m[string(t[i])]; ok {
// 			if val != string(s[i]) {
// 				status = false
// 				break
// 			}
// 		} else {
// 			m[string(t[i])] = string(s[i])
// 		}
// 		i++
// 	}

// 	fmt.Println(m)
// 	fmt.Println(status)
// 	return status
// }
