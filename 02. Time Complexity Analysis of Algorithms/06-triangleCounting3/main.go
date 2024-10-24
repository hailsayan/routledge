package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	counter := 0

	for i := 1; i <= n/3; i++ {
		counter += ((n - 3*i) / 2) - int(math.Max(0, float64((n/2)-(2*i)+1))) + 1
	}

	fmt.Println(counter)
}
