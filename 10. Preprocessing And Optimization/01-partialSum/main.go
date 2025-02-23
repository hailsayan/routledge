package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var n, q int
	scanf("%d %d\n", &q, &n)

	sequences := make([]int, q)
	ps := make([]int, q+1)

	for i := 0; i < q; i++ {
		scanf("%d ", &sequences[i])
	}

	ps[0] = 0
	for i := 1; i < q+1; i++ {
		ps[i] = ps[i-1] + sequences[i-1]
	}

	var l, r int
	for i := 0; i < n; i++ {
		scanf("%d %d\n", &l, &r)
		fmt.Println(ps[r+1] - ps[l])
	}
}
