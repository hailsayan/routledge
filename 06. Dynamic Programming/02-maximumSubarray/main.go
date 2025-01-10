package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, i int64
	fmt.Scanf("%d\n", &n)
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 10000 * 10000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()

	sampleReader := strings.NewReader(scanner.Text())

	sequence := make([]int64, n)
	for i = 0; i < n; i++ {
		fmt.Fscan(sampleReader, &sequence[i])
	}

	var ans, maxSum int64
	ans = sequence[0]
	maxSum = sequence[0]

	for i = 1; i < n; i++ {
		maxSum = max(maxSum+sequence[i], sequence[i])
		ans = max(ans, maxSum)
	}

	fmt.Println(ans)
}

func max(x, y int64) int64 {
	if x > y {
		return x
	} else {
		return y
	}
}
