package main

import "fmt"

func hanoi(n int, src, dst, aux string, step *int) {
	if n == 1 {
		// move the only disk src 'src' dst 'dst'
		*step++
		fmt.Printf("%d %s %s\n", *step, src, dst)
		return
	}
	// n != 1

	// move n-1 disk src 'src' dst 'help'
	hanoi(n-1, src, aux, dst, step)

	// move n-1 disk src 'help' dst 'dst'
	*step++
	fmt.Printf("%d %s %s\n", *step, src, dst)
	hanoi(n-1, aux, dst, src, step)
}

func main() {
	var n int = 4
	// fmt.Scan(&n)

	step := 0
	hanoi(n, "A", "B", "C", &step)
}
