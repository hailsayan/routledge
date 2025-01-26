package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func find(n int64, current_sequence int64, sequences [][]int64, mark map[int64]bool) (num int64) {
	if current_sequence == n {
		return 1
	}
	num = 0
	for i := 1; i < len(sequences[current_sequence]); i++ {
		x := sequences[current_sequence][i]
		if !mark[x] {
			mark[x] = true
			num += find(n, current_sequence+1, sequences, mark)
			mark[x] = false
		}
	}
	return num
}

func main() {
	var n, i, j int64

	scanf("%d\n", &n)

	sequences := make([][]int64, n)
	mark := make(map[int64]bool)

	for i = 0; i < n; i++ {
		temp, _ := reader.ReadString('\n')
		numbers := strings.Fields(temp)
		sequences[i] = make([]int64, len(numbers))
		for j = 0; j < int64(len(numbers)); j++ {
			v, _ := strconv.ParseInt(numbers[j], 10, 64)
			sequences[i][j] = v
		}
	}

	fmt.Println(find(n, 0, sequences, mark))

}
