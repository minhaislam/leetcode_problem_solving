package main

import "fmt"

func main() {

	haystack := "sadbutsad"
	needle := "sad"

	// output := []int{1, 2, 2, 3, 5, 6}

	strStr(haystack, needle)

}

func strStr(haystack string, needle string) int {
	// range_val := len(needle)
	i := 0
	occurance := -1
	for i+len(needle) <= len(haystack) && i < len(haystack) {
		if haystack[i:i+len(needle)] == needle {
			occurance = i
			// i = i + range_val
			break
		} else {
			i++
		}
	}

	fmt.Println(occurance)
	return occurance
}
