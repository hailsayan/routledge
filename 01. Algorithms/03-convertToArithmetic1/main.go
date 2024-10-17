package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	m, M := a[0], a[0]
	for _, num := range a {
		if num < m {
			m = num
		}
		if num > M {
			M = num
		}
	}

	minX := m - (n-1)*k
	maxX := M

	minCost := math.MaxInt64

	for x := minX; x <= maxX; x++ {
		cost := 0
		for i := 0; i < n; i++ {
			expected := x + i*k
			cost += abs(expected - a[i])
		}
		if cost < minCost {
			minCost = cost
		}
	}

	fmt.Println(minCost)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
