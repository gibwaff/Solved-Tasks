package main

import "fmt"

func FindUniq(arr []float32) float32 {
	var res float32
	if arr[0] == arr[1] {
		for i := 2; i < len(arr); i++ {
			if arr[i] != arr[0] {
				res = arr[i]
				break
			}
		}
	} else {
		if arr[2] == arr[0] {
			res = arr[1]
		} else {
			res = arr[0]
		}
	}
	return res
}

func main() {
	fmt.Println(FindUniq([]float32{1, 1, 1, 2, 1, 1})) // Output: 2
}
