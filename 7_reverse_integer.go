package main

import (
	"fmt"
	"strconv"
)

func main() {
	var val int = 214748364
	final_val := reverse(val)
	fmt.Println(final_val)

}
func reverse(x int) int {
	// string_val := strconv.FormatInt(int64(x), 10)
	string_val_re := "0"
	final := 0
	if x >= -2147483648 && x <= 2147483647 {

		string_val_re = strconv.FormatInt(int64(x), 10)
		final = reverse_fuc(string_val_re)
	}
	return final

}

func reverse_fuc(string_val string) int {
	// string_val := strconv.FormatInt(int64(x), 10)
	// fmt.Println(string_val)
	// fmt.Printf("type %T\n", string_val)
	i := len(string_val)
	reverse_str := ""
	for i > 0 {
		// fmt.Println(string(string_val[i-1]))
		if string(string_val[i-1]) == "-" {
			reverse_str = string(string_val[i-1]) + reverse_str
		} else {
			reverse_str = reverse_str + string(string_val[i-1])
		}
		i--

	}
	// fmt.Println(reverse_str)

	reverse_str2, err := strconv.ParseInt(reverse_str, 10, 32)
	if err != nil {
		return 0
	} else {
		return int(reverse_str2)
	}
}
