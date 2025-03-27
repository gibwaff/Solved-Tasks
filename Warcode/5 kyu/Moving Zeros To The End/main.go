package main

import "fmt"

func MoveZeros(arr []int) []int {
	cnt := 0
	for _, el := range arr {
		if el == 0 {
			cnt++
		}
	}
	ResArray := make([]int, len(arr))
	res_index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			ResArray[res_index] = arr[i]
			res_index++
		}
	}
	for ; res_index < len(ResArray); res_index++ {
		ResArray[res_index] = 0
	}

	return ResArray

}

func main() {
	arr := []int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}
	fmt.Println(MoveZeros(arr))
}
