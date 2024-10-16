package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstLine, _ := reader.ReadString('\n')
	firstLine = strings.TrimSpace(firstLine)
	parts := strings.Split(firstLine, " ")
	n, _ := strconv.Atoi(parts[0])
	q, _ := strconv.Atoi(parts[1])

	secondLine, _ := reader.ReadString('\n')
	secondLine = strings.TrimSpace(secondLine)
	sequence := make([]int, n)
	seqParts := strings.Split(secondLine, " ")
	for i := 0; i < n; i++ {
		sequence[i], _ = strconv.Atoi(seqParts[i])
	}

	results := make([]int, q)
	for i := 0; i < q; i++ {
		queryLine, _ := reader.ReadString('\n')
		queryLine = strings.TrimSpace(queryLine)
		x, _ := strconv.Atoi(queryLine)

		count := 0
		for _, num := range sequence {
			if num < x {
				count++
			}
		}
		results[i] = count
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
