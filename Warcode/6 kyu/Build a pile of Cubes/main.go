package main

import "fmt"

func FindNb(m int) int {
	res := 0
	for i := 1; res < m; i++ {
		res += i * i * i
		if res == m {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(FindNb(1071225)) // 45
}
