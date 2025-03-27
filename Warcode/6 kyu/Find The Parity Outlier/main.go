package main

import "fmt"

func FindOutlier(integers []int) (res int) {
	m0 := integers[0] % 2
	m1 := integers[1] % 2
	if m0 == m1 {
		if m0 == 0 {
			for _, el := range integers {
				if el%2 != 0 {
					res = el
					break
				}
			}
		} else {
			for _, el := range integers {
				if el%2 == 0 {
					res = el
					break
				}
			}
		}
	} else {
		m2 := integers[2] % 2
		if m2 == m0 {
			res = integers[1]
		} else {
			res = integers[0]
		}
	}

	return
}

func main() {
	fmt.Println(FindOutlier([]int{2, 6, 8, -10, 3})) // 3
}
