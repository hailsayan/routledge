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

	var n, S, tmp int
	scanf("%d %d\n", &n, &S)

	if n == 0 {
		printf("%d", 0)
		writer.Flush()
		return
	}

	prices := make([][3]int, n+1)
	for i := 0; i < n; i++ {
		scanf("%d ", &tmp)
		prices[i][0] = tmp
		min := 2000
		total := 0
		for j := 0; j < prices[i][0]; j++ {
			scanf("%d ", &tmp)
			total += tmp
			if tmp < min {
				min = tmp
			}
		}
		prices[i][1] = min
		prices[i][2] = total
	}
	var s [2001][2001]int
	var pr [2001][2001]byte

	for j := 1; j <= S; j++ {
		if prices[0][2] <= j {
			s[0][j] = prices[0][0]
			pr[0][j] = '2'
		} else if prices[0][1] <= j {
			s[0][j] = 1
			pr[0][j] = '1'
		} else {
			s[0][j] = 0
			pr[0][j] = '0'
		}
	}
	for i := 0; i < n; i++ {
		s[i][0] = 0
		pr[i][0] = '0'
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= S; j++ {
			a2 := 0
			if j >= prices[i][2] {
				a2 = s[i-1][j-prices[i][2]] + prices[i][0]
			}
			a1 := 0
			if j >= prices[i][1] {
				a1 = s[i-1][j-prices[i][1]] + 1
			}
			a0 := s[i-1][j]

			s[i][j] = max(a0, a1, a2)
			if s[i][j] == a0 {
				pr[i][j] = '0'
			} else if s[i][j] == a1 {
				pr[i][j] = '1'
			} else {
				pr[i][j] = '2'
			}
		}
	}

	printf("%d\n", s[n-1][S])

	res := make([]byte, n)
	for i := n - 1; i >= 0; i-- {
		res[i] = pr[i][S]
		if pr[i][S] == '2' {
			S -= prices[i][2]
		} else if pr[i][S] == '1' {
			S -= prices[i][1]
		}
	}
	printf("%s", res)
	writer.Flush()
}

func max(a ...int) int {
	max := 0
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}
