package main

import "fmt"

func fn(x int) int {
	if x == 0 {
		return 5
	}
	tmp := fn(x-1)
	if x%2 == 0{
		return tmp - 21
	} else{
		return tmp * tmp
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println(fn(n))
}
