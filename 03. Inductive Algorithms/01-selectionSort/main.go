package main

import (
	"fmt"
	"strings"
)

func selectionSort(arr []int, n int) []int {
	for j := 0; j < n; j++ {
		minIdx := j
		for i := j + 1; i < n; i++ {
			if arr[i] < arr[minIdx] {
				minIdx = i
			}
		}
		arr[minIdx], arr[j] = arr[j], arr[minIdx]
	}
	return arr
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(strings.Trim(fmt.Sprint(selectionSort(arr, n)), "[]"))
}
