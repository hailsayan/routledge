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

func main() {

	var n, m int
	scanf("%d %d\n", &n, &m)
	a := make([][]int, n)

	var tmp int
	for r := 0; r < n; r++ {
		row := make([]int, m)
		for c := 0; c < m; c++ {
			scanf("%d ", &tmp)
			row[c] = tmp
		}
		a[r] = row
	}

	max := a[0][0]
	for r1 := 0; r1 < n; r1++ {
		for r2 := r1; r2 < n; r2++ {
			b := make([]int, m)
			for c := 0; c < m; c++ {
				sum := 0
				for r := r1; r <= r2; r++ {
					sum += a[r][c]
				}
				b[c] = sum
				if c > 0 {
					if b[c-1]+b[c] > b[c] {
						b[c] = b[c-1] + b[c]
					}
				}
				if b[c] > max {
					max = b[c]
				}
			}

		}
	}

	printf("%d", max)

	writer.Flush()
}
