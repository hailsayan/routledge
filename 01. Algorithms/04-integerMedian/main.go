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

	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	aStr, _ := reader.ReadString('\n')
	aStrSlice := strings.Fields(strings.TrimSpace(aStr))
	a := make([]int, n)

	for i, val := range aStrSlice {
		a[i], _ = strconv.Atoi(val)
	}

	sort.Ints(a)

	var median int
	if n%2 == 1 {
		median = a[n/2]
	} else {
		median = a[n/2-1]
	}

	totalCost := 0
	for _, val := range a {
		totalCost += abs(val - median)
	}

	fmt.Printf("%d %d\n", median, totalCost)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
