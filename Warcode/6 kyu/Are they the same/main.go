package main

import "fmt"

func sort(a []int) []int { //bubble sort
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}

func compSame(a, b []int) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	a, b = sort(a), sort(b)
	for i, j := 0, 0; i < len(b); {
		if j > len(a) {
			return false
		}
		if b[i] < a[j]*a[j] {
			return false
		}
		if b[i] > a[j]*a[j] {
			j++
			continue
		}
		i++
	}

	return true
}

func main() {
	a, b := []int{121, 144, 19, 161, 19, 144, 19, 11}, []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(compSame(a, b))
}
