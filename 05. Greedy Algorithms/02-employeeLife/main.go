package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func maxTasks(deadlines []int) int {
	sort.Ints(deadlines)

	time, count := 0, 0

	for _, deadline := range deadlines {
		if time+1 <= deadline {
			time++
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	nLine, _ := reader.ReadString('\n')
	nLine = strings.TrimSpace(nLine)
	n, _ := strconv.Atoi(nLine)

	deadlinesLine, _ := reader.ReadString('\n')
	deadlinesLine = strings.TrimSpace(deadlinesLine)
	deadlineStrs := strings.Split(deadlinesLine, " ")

	deadlines := make([]int, n)
	for i, d := range deadlineStrs {
		deadlines[i], _ = strconv.Atoi(d)
	}

	fmt.Println(maxTasks(deadlines))
}
