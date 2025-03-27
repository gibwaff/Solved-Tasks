package main

import "fmt"

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	cur := 1
	pos, peaks := []int{}, []int{}
	for i := 2; i < len(array); i++ {
		if array[i] > array[cur] || array[i] > array[i-1] {
			cur = i
		} else if array[i] < array[cur] && array[cur-1] < array[cur] {
			pos = append(pos, cur)
			peaks = append(peaks, array[cur])
			cur++
		}
	}

	return PosPeaks{Pos: pos, Peaks: peaks}
}

func main() {
	arr := []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3} //{[3 7] [6 3]}
	fmt.Println(PickPeaks(arr))
}
