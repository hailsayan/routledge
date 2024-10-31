package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, x, i int64
	fmt.Scanf("%d %d\n", &n, &x)
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 10000 * 10000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()

	sampleReader := strings.NewReader(scanner.Text())

	sequence := make([]int64, n+1)
	for i = 0; i < n+1; i++ {
		fmt.Fscan(sampleReader, &sequence[i])
	}

	var p int64
	p = sequence[0]
	for i = 1; i <= n; i++ {
		p = (p*x + sequence[i]) % 1000000007
		if p < 0 {
			p += 1000000007
		}
	}

	fmt.Println(p)
}
