package main

import (
	"fmt"
	"strings"
)


func bubbleSort(a []int, n int) []int {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func main() {
	var n int // = 11
	//arr := []int{2, 10, 7, 12, 3, 1, 4, 5, 2, 2, 9}
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(strings.Trim(fmt.Sprint(bubbleSort(arr, n)), "[]"))
}
