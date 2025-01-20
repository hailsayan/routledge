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
	var P, S string
	scanf("%s\n%s", &P, &S)

	var dp [2001][2001]int
	var pr [2001][2001]int

	for i := 1; i < len(P)+1; i++ {
		for j := 1; j < len(S)+1; j++ {
			if P[i-1] == S[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				pr[i][j] = 0
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
					pr[i][j] = 1
				} else {
					dp[i][j] = dp[i][j-1]
					pr[i][j] = 2
				}
			}
		}
	}

	i := len(P)
	j := len(S)
	index := dp[len(P)][len(S)] - 1
	res := make([]byte, index+1)
	for index >= 0 {
		switch pr[i][j] {
		case 0:
			res[index] = P[i-1]
			i--
			j--
			index--
		case 1:
			i--
		case 2:
			j--
		}

	}

	printf("%d\n%s", dp[len(P)][len(S)], res)
	writer.Flush()
}
