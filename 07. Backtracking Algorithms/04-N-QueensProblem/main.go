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
var table [8][8]int

func main() {

	scanf("%d %d\n", &n, &k)

	sum(0, 0)

	printf("%d", total)

	writer.Flush()
}

func sum(q, h int) {
	if q == k {
		total++
		return
	}
	if h == n*n {
		return
	}

	if check(h) {
		r := h / n
		c := h % n
		table[r][c] = 1
		sum(q+1, h+1)
		table[r][c] = 0
	}
	sum(q, h+1)
}

func check(h int) bool {
	for i := 0; i < h; i++ {
		r := h / n
		c := h % n
		rr := i / n
		cc := i % n
		if table[rr][cc] == 1 {
			rrr := rr - r
			ccc := cc - c
			if rrr == ccc || rrr == -ccc || rrr == 0 || ccc == 0 {
				return false
			}
		}

	}
	return true
}
