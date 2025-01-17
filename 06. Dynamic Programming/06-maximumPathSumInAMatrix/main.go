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
	var n, m, tmp int
	scanf("%d %d\n", &n, &m)
	A := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, m)
		for j := 0; j < m; j++ {
			scanf("%d ", &tmp)
			row[j] = tmp
		}
		A[i] = row
	}

	s := make([][]int, n)
	p := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		s[i] = make([]int, m)
		p[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if j == 0 && i == n-1 {
				s[i][j] = A[i][j]
				p[i][j] = 0
			} else if j == 0 {
				// the path from lower cell is better
				s[i][j] = s[i+1][j] + A[i][j]
				p[i][j] = 1
			} else if i == n-1 {
				s[i][j] = s[i][j-1] + A[i][j]
				p[i][j] = 2
			} else if s[i+1][j] > s[i][j-1] {
				// the path from lower cell is better
				s[i][j] = s[i+1][j] + A[i][j]
				p[i][j] = 1
			} else {
				// the path from the left cell is better
				s[i][j] = s[i][j-1] + A[i][j]
				p[i][j] = 2
			}
		}
	}

	ans := make([]byte, n+m-2)
	index := n + m - 3
	currentRow := 0
	currentColumn := m - 1
	for !(currentRow == n-1 && currentColumn == 0) {
		// printf("%d %d\n", currentRow, currentColumn)
		if p[currentRow][currentColumn] == 1 {
			ans[index] = 'U'
			// printf("U")
			currentRow++
		} else {
			ans[index] = 'R'
			// printf("R")
			// ans = "R" + ans
			currentColumn--
		}
		index--
		// printf("\n%d %d", currentRow, currentColumn)
	}
	printf("%d\n%s", s[0][m-1], ans)
	writer.Flush()
}
