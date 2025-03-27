package main

import "fmt"

func Solution(str string) []string {
	var length int
	if len(str)%2 == 0 {
		length = len(str) / 2
	} else {
		length = (len(str) + 1) / 2
	}

	rune_str := []rune(str)
	res := make([]string, length)
	for i, j := 0, 0; i < len(str); i, j = i+2, j+1 {
		if i+1 != len(str) {
			res[j] = string(rune_str[i]) + string(rune_str[i+1])
		} else {
			res[j] = string(rune_str[i]) + "_"
		}
	}

	return res

}

func main() {
	fmt.Println(Solution("abcde")) //  ["ab", "cd", "e_"]
}
