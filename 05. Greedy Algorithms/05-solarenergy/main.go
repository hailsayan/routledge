package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

type key struct {
	id       int64
	producer int64
}

func main() {
	var n, k, i int64
	scanf("%d %d\n", &n, &k)

	mapper := make(map[key]int64)

	var consume, energy int64
	for i = 0; i < n; i++ {
		scanf("%d %d\n", &consume, &energy)
		pr := key{id: i, producer: energy}
		mapper[pr] = consume
	}

	keys := make([]key, 0, len(mapper))

	for ke := range mapper {
		keys = append(keys, ke)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return mapper[keys[i]] < mapper[keys[j]]
	})

	var sum int64
	sum = k

	for _, k2 := range keys {
		if mapper[k2] <= sum {
			if k2.producer > mapper[k2] {
				sum = sum + k2.producer - mapper[k2]
			}
		}
	}

	fmt.Println(sum)
}
