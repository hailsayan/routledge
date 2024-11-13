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
	secondLine, _ := reader.ReadString('\n')

	firstLine = strings.TrimSpace(firstLine)
	params := strings.Split(firstLine, " ")
	n, _ := strconv.Atoi(params[0])
	r, _ := strconv.ParseInt(params[1], 10, 64)

	secondLine = strings.TrimSpace(secondLine)
	priceStrings := strings.Split(secondLine, " ")
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		prices[i], _ = strconv.Atoi(priceStrings[i])
	}

	sort.Ints(prices)

	count := 0
	for _, price := range prices {
		if r >= int64(price) {
			r -= int64(price)
			count++
		} else {
			break
		}
	}
	fmt.Println(count)
}
