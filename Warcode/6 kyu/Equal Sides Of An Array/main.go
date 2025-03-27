package main

import "fmt"

func FindEvenIndexx(nums []int) int {
	cur, left_s, right_s := 0, 0, 0
	for i := 1; i < len(nums); i++ {
		right_s += nums[i]
	}
	for ; cur < len(nums)-1; cur++ {
		if left_s == right_s {
			return cur
		}
		left_s += nums[cur]
		right_s -= nums[cur+1]
	}
	if left_s == right_s {
		return cur
	}
	return -1

}

func main() {
	arr := []int{1, 100, 50, -51, 1, 1}
	fmt.Println(FindEvenIndexx(arr)) // 1
}
