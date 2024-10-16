package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, q, i int64
	fmt.Scanf("%d %d\n", &n, &q)
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 1000 * 1000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()

	sampleReader := strings.NewReader(scanner.Text())

	sequence := make([]int64, n)
	for i = 0; i < n; i++ {
		fmt.Fscan(sampleReader, &sequence[i])
	}

	m := max(sequence) + 1
	ct := make([]int64, m)
	for i = 0; i < n; i++ {
		ct[sequence[i]]++
	}

	ps := make([]int64, m)

	for i = 1; i < m; i++ {
		ps[i] = ct[i] + ps[i-1]
	}

	result := make([]int64, q)
	var iter int64
	for i = 0; i < q; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &iter)
		if iter >= m {
			result[i] = ps[m-1]
		} else {
			result[i] = ps[iter] - ct[iter]
		}
	}

	for i = 0; i < q; i++ {
		fmt.Println(result[i])
	}
}

func max(array []int64) (max int64) {
	max = array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}
