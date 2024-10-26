package main

import (
	"fmt"
	"strings"
)

func insertionSort(a []int, n int) []int {
	for i := 0; i < n; i++ {
		p := i
		item := a[p]
		for 0 < p && a[p-1] > item {
			a[p] = a[p-1]
			p = p-1
		}
		a[p] = item
	}
	return a
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(strings.Trim(fmt.Sprint(insertionSort(arr, n)), "[]"))
}
