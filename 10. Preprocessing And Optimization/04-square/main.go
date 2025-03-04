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
	var n, m, k, tmp, total int
	scanf("%d %d %d ", &n, &m, &k)
	a := make([][]int, n)
	dp := make([][]int, n+1)

	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			scanf("%d ", &tmp)
			a[i][j] = tmp
			if tmp == 0 {
				total++
			}

		}
	}

	// printf("%v\n", a)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			dp[i][j] = a[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]

		}
	}
	// printf("%v\n", dp)
	max := 0
	for i := 0; i <= n-k; i++ {
		for j := 0; j <= m-k; j++ {
			X := i + k - 1
			Y := j + k - 1
			x := i
			y := j
			a := dp[X+1][Y+1] - dp[X+1][y] - dp[x][Y+1] + dp[x][y]
			// printf("%d ", a)
			for ii := 0; ii <= n-k; ii++ {
				for jj := 0; jj <= m-k; jj++ {
					// printf("%d %d %d %d ", i, j, ii, jj)
					// writer.Flush()
					// if !(ii == i && jj == j) {

					XX := ii + k - 1
					YY := jj + k - 1
					xx := ii
					yy := jj

					// printf(" x:%d  y:%d  X:%d  Y:%d\n", x, y, X, Y)
					// printf("xx:%d yy:%d XX:%d YY:%d\n", xx, yy, XX, YY)

					// writer.Flush()
					b := dp[XX+1][YY+1] - dp[XX+1][yy] - dp[xx][YY+1] + dp[xx][yy]

					sum := a + b
					c := 0

					cx := xx
					cy := yy
					cX := X
					cY := Y

					if x > xx {
						cx = x
					}
					if y > yy {
						cy = y
					}
					if XX < X {
						cX = XX
					}
					if YY < Y {
						cY = YY
					}
					// printf("cx:%d cy:%d cX:%d cY:%d\n", cx, cy, cX, cY)

					if cX-cx >= 0 && cY-cy >= 0 {
						// printf("C:= %d %d %d %d\n", xx, yy, X, Y)
						c = dp[cX+1][cY+1] - dp[cX+1][cy] - dp[cx][cY+1] + dp[cx][cy]
						sum -= c
					}
					// printf("%d\n", sum)
					// printf("%d %d+%d-%d=%d\n", total, a, b, c, sum)
					// printf("--------------\n")

					if sum > max {
						// printf("%d\n", max+total)
						// printf("%d\n", dp[X+1][Y+1]-dp[X+1][yy]-dp[xx][Y+1]+dp[xx][yy])
						max = sum
						// printf("%d %d\n", i, j)
						// printf("%d %d+%d-%d=%d\n", total, a, b, c, sum)
						// printf("%d %d %d %d\n", x, y, xx, yy)
					}

					// }
				}
			}
		}
	}
	printf("%d", max+total)
	// for i := 0; i < q; i++ {
	// 	scanf("%d %d %d %d ", &x, &y, &X, &Y)
	// 	sum := dp[X+1][Y+1] - dp[X+1][y] - dp[x][Y+1] + dp[x][y]
	// 	printf("%d\n", sum)

	// }
	writer.Flush()
}

// func merge(aa *[]int, bb *[]int) {
// 	a := *aa
// 	b := *bb
// 	if len(b) == 0 {
// 		return
// 	}
// 	if len(a) == 0 {
// 		*aa, *bb = *bb, *aa
// 		return
// 	}
// 	ans := make([]int, len(a)+len(b))
// 	l := 0
// 	r := 0
// 	for i := 0; i < len(a)+len(b); i++ {
// 		if l == len(a) {
// 			ans[i] = b[r]
// 			r++
// 		} else if r == len(b) {
// 			ans[i] = a[l]
// 			l++
// 		} else if a[l] > b[r] {
// 			ans[i] = b[r]
// 			r++
// 		} else {
// 			ans[i] = a[l]
// 			l++
// 		}
// 	}
// 	*aa = ans
// 	*bb = make([]int, 0)
// }
