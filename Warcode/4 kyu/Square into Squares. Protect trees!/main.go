package main

import "fmt"

func Decompose(n int64) []int64 {
	// your code
}

func NearestSqrt(n int64) (res int64) {
	for i := 1; i*i <= int(n); i++ {
		res = int64(i * i)
	}
	return
}

func main() {
	fmt.Println(Decompose(11)) // should return [1, 2, 4, 10]
	fmt.Println(Decompose(44)) // should return [2, 3, 5, 7, 43]
}
