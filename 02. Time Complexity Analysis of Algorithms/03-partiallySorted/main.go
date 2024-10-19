package main

import (
	"fmt"
)

func main() {
	var n int // = 5
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	count := 0
	for i := 0; i < n; i++ {
		if arr[i] != i+1 {
			count = count + 1
		}
	}
	if count == 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
