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

	var n, tmp, ans int
	scanf("%d\n", &n)
	var h [309]int
	var dp [309][309]int
	mod := int(1e9 + 7)
	for i := 1; i <= n; i++ {
		scanf("%d ", &tmp)
		h[i] = tmp
	}

	for i := 1; i <= n; i++ {
		dp[i][i+1] = 1
	}
	ans++
	for length := 2; length <= n; length++ {
		for l := 1; l <= n-length+1; l++ {
			r := l + length
			dp[l][r] = dp[l][r-1]
			if h[l] == h[r-1] {
				dp[l][r]++

				for x := l + 1; x <= r-2; x++ {
					if h[x] >= h[l] {
						dp[l][r] += dp[x][r-1]
						dp[l][r] = ((dp[l][r] % mod) + mod) % mod
					}
				}
			}
		}
		ans += dp[n-length+1][n+1]

	}

	ans = ((ans % mod) + mod) % mod

	printf("%d\n", ans)
	writer.Flush()
}
