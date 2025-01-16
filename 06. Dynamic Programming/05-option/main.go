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

// combination
type com struct {
	r int
	n int
}

func main() {
	var q int
	scanf("%d\n", &q)
	A := make([]com, q)

	var tmp1, tmp2 int
	for i := 0; i < q; i++ {
		scanf("%d %d\n", &tmp1, &tmp2)
		A[i].n = tmp1
		A[i].r = tmp2
	}
	mod := int(1e9 + 7)
	s := make([][]int, 2001)
	for i := 0; i < 2001; i++ {
		s[i] = make([]int, 2001)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				s[i][j] = 1
			} else {
				s[i][j] = (s[i-1][j] + s[i-1][j-1]) % mod
			}
		}
	}
	for _, a := range A {
		printf("%d\n", s[a.n][a.r])
	}
	writer.Flush()
}
