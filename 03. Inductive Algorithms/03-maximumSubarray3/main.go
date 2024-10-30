package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func Scan(a ...any)  { fmt.Fscan(reader, a...) }
func Print(a ...any) { fmt.Fprint(writer, a...) }

func main() {
	defer writer.Flush()

	var n int
	Scan(&n)

	ns := make([]int, n)
	var in int
	for i := range ns {
		Scan(&in)
		ns[i] = in
	}

	ans := -99999999999
	maxi := 0

	for i := range ns {
		maxi = max(maxi+ns[i], ns[i])
		ans = max(ans, maxi)
	}

	Print(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
