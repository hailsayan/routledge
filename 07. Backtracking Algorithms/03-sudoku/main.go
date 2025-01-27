package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriterSize(os.Stdout, math.MaxInt32)

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

var total int

var ans [9][9]int

func main() {

	var tmp int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			scanf("%d ", &tmp)
			ans[i][j] = tmp
		}
	}

	sum(0)
	if total == 0 {
		printf("%s", "No solution exists")
	}
	writer.Flush()
}

func sum(n int) bool {

	if n == 81 {
		total++
		print()
		return true
	}

	r := n / 9
	c := n % 9

	if ans[r][c] != 0 {
		if sum(n + 1) {
			return true
		}
		return false
	}

	for j := 1; j < 10; j++ {
		if check(r, c, j) {
			ans[r][c] = j
			if sum(n + 1) {
				return true
			}
			ans[r][c] = 0

		}
	}
	return false
}
func print() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			printf("%d ", ans[i][j])
		}
		printf("\n")
	}
	writer.Flush()
}
func check(r, c, n int) bool {
	for i := 0; i < 9; i++ {
		if ans[r][i] == n {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if ans[i][c] == n {
			return false
		}
	}

	for i := r - (r % 3); i < r-(r%3)+3; i++ {
		for j := c - (c % 3); j < c-(c%3)+3; j++ {
			if ans[i][j] == n {
				return false
			}
		}
	}
	return true
}
