package main

import "fmt"

func SOD (n int) int {
	if n == 0 {
		return 0
	}
	return n % 10 + SOD(n/10)
}

func main(){
	var n int
	fmt.Scan(&n)

	fmt.Println(SOD(n))
}