package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var n, c int
	scanf("%d %d\n", &n, &c)

	sequence := make([]int, n)

	for i := 0; i < n; i++ {
		scanf("%d", &sequence[i])
	}

	sort.Ints(sequence)

	for i := n - 1; i >= 0; i-- {
		c = c - int(math.Min(float64(c), math.Max(0, float64(sequence[i]-c))))
	}

	fmt.Println(c)

}
