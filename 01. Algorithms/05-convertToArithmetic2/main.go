package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstLine, _ := reader.ReadString('\n')
	firstLineSplit := strings.Fields(strings.TrimSpace(firstLine))
	n, _ := strconv.Atoi(firstLineSplit[0])
	k, _ := strconv.Atoi(firstLineSplit[1])

	secondLine, _ := reader.ReadString('\n')
	secondLineSplit := strings.Fields(strings.TrimSpace(secondLine))
	a := make([]int, n)
	for i, val := range secondLineSplit {
		a[i], _ = strconv.Atoi(val)
	}

	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = a[i] - i*k
	}

	sort.Ints(b)
	median := b[n/2]

	operations := 0
	for _, val := range b {
		operations += abs(val - median)
	}

	fmt.Println(operations)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
