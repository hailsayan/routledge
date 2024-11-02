package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	fmt.Print(gcd(x, y))
}
