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
	var n, q, tmp, y, x, Y, X int
	scanf("%d %d\n", &n, &q)
	m := make([][]int, n)
	dp := make([][]int, n+1)

	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			scanf("%d ", &tmp)
			m[i][j] = tmp
		}
	}

	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			dp[i][j] = m[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]

		}
	}
	for i := 0; i < q; i++ {
		scanf("%d %d %d %d ", &x, &y, &X, &Y)
		sum := dp[X+1][Y+1] - dp[X+1][y] - dp[x][Y+1] + dp[x][y]
		printf("%d\n", sum)

	}
	writer.Flush()
}
