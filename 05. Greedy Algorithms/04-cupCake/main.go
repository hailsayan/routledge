package main

import (
	"fmt"
	"os"
)

const maxn = 5000 + 2

func main() {
	var n, k int
	var c [maxn]int
	var money int

	fmt.Fscan(os.Stdin, &n, &k)
	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &c[i])
	}

	if k >= 3 {
		money = c[0]
		for i := 0; i < n; i++ {
			money = min(money, c[i])
		}
	} else if k == 2 {
		var maxs [maxn]int
		var maxe [maxn]int

		maxs[0] = c[0]
		for i := 1; i < n; i++ {
			maxs[i] = max(c[i], maxs[i-1])
		}

		maxe[n-1] = c[n-1]
		for i := n - 2; i >= 0; i-- {
			maxe[i] = max(c[i], maxe[i+1])
		}

		money = 5000
		for i := 1; i < n; i++ {
			money = min(money, min(maxs[i-1], maxe[i]))
		}
	} else {
		money = c[0]
		for i := 0; i < n; i++ {
			money = max(money, c[i])
		}
	}

	fmt.Println(money)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
