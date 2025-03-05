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

const ln = 100001

func main() {
	var t, k, a, b, count int
	scanf("%d %d\n", &t, &k)
	dp := make([]int, ln)
	sp := make([]int, ln)

	// scanf("%d ", &tmp)

	for i := 0; i < ln; i++ {
		if i-k < 0 {
			dp[i] = 1
			continue
		}
		dp[i] = (dp[i-1] + dp[i-k]) % int(1e9+7)
	}
	sp[0] = dp[0]
	for i := 1; i < ln; i++ {
		sp[i] = (sp[i-1] + dp[i]) % int(1e9+7)
	}

	for i := 0; i < t; i++ {
		scanf("%d %d\n", &a, &b)
		if a == 0 {
			count = sp[b]
		} else {
			count = (sp[b] - sp[a-1] + int(1e9+7)) % int(1e9+7)
		}
		printf("%d\n", count)
	}

	writer.Flush()
}
