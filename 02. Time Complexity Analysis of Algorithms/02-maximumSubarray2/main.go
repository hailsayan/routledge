package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func kadane(arr []int) int {
	currentSum := math.MinInt32
	maxSum := math.MinInt32

	for _, num := range arr {
		currentSum = max(currentSum+num, num)
		maxSum = max(currentSum, maxSum)
	}
	return maxSum

}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// arr := [6]int{-7, 3, -1, 2, -4, 3}
	// fmt.Print(kadane(arr))
	fmt.Println((kadane(arr)))
}
