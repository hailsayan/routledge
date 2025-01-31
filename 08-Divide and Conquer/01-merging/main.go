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

	var tmp, n, k int

	scanf("%d %d\n", &n, &k)
	a := make([][]int, k)

	for i := 0; i < k; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			scanf("%d ", &tmp)
			a[i][j] = tmp
		}
	}
	answer := mergeSort(a, 0, k)

	for j := 0; j < n*k; j++ {
		printf("%d ", answer[j])
	}

	writer.Flush()
}
func mergeSort(a [][]int, l, r int) []int {
	if r-l == 1 {
		return a[l]
	}
	if r-l == 0 {
		return make([]int, 0)
	}

	mid := (r + l) / 2
	right := mergeSort(a, r, mid)
	left := mergeSort(a, mid, l)

	return merge(right, left)
}

func merge(f []int, s []int) []int {
	r := 0
	l := 0
	ans := make([]int, len(f)+len(s))
	for i := 0; i < len(f)+len(s); i++ {
		if r >= len(f) {
			ans[i] = s[l]
			l++
		} else if l >= len(s) {
			ans[i] = f[r]
			r++
		} else if f[r] > s[l] {
			ans[i] = s[l]
			l++
		} else {
			ans[i] = f[r]
			r++
		}
	}
	return ans
}