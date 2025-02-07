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

var ans int

func main() {
	var n, tmp int
	scanf("%d\n", &n)
	nn := int(math.Pow(2, float64(n)))
	a := make([]int, nn)
	for i := 0; i < nn; i++ {
		scanf("%d ", &tmp)
		a[i] = tmp
	}
	bianry(a, 0)
	printf("%d\n", ans)
	writer.Flush()

}
func bianry(a []int, sum int) {
	l := len(a)
	if len(a) == 1 {
		s := sum + a[0]
		if s > ans {
			ans = s
		}
		return
	}
	h := int(uint(l) >> 1)
	bianry(a[0:h], max(a[h:])+sum)
	bianry(a[h:], max(a[0:h])+sum)
}

func max(a []int) int {
	m := 0
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}