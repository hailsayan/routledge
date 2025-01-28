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

var total, n, k int
var table [10][10]int

func main() {

	scanf("%d %d\n", &n, &k)

	sum(0, 0, 0)
	print()
	printf("%d", total)

	writer.Flush()
}

func sum(q, r, c int) {

	if q == k {
		table[r][c] = 1
		return
	}
	rr := r
	cc := c
	for i := 1; i <= 8; i++ {
		switch i {
		case 1:
			r -= 2
			c--
		case 2:
			r -= 2
			c++
		case 3:
			r--
			c += 2
		case 4:
			r++
			c += 2
		case 5:
			r += 2
			c++
		case 6:
			r += 2
			c--
		case 7:
			r++
			c -= 2
		case 8:
			r--
			c -= 2
		}
		if r >= 0 && r < n && c >= 0 && c < n {
			sum(q+1, r, c)
		}

		r = rr
		c = cc
	}
}

func print() {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if table[i][j] == 1 {
				total++
			}
		}
	}
}