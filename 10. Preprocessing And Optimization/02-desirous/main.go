package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var n, q int
	scanf("%d %d\n", &n, &q)

	sets := make(map[int]map[int]bool)
	for i := 1; i <= n; i++ {
		sets[i] = make(map[int]bool)
		sets[i][i] = true
	}

	for q > 0 {
		var t int
		scanf("%d ", &t)

		if t == 1 {
			var a, b int
			scanf("%d %d\n", &a, &b)
			for k := range sets[a] {
				sets[b][k] = true
			}
			sets[a] = make(map[int]bool)
		} else if t == 2 {
			var c int
			scanf("%d\n", &c)
			fmt.Println(len(sets[c]))
		} else if t == 3 {
			var d int
			scanf("%d\n", &d)
			var arr []int
			for k := range sets[d] {
				arr = append(arr, k)
			}
			sort.Ints(arr)
			if len(arr) == 0 {
				fmt.Println(-1)
			} else {
				for _, num := range arr {
					fmt.Printf("%d ", num)
				}
				fmt.Println()
			}
		}

		q--
	}
}
