package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) {
	fmt.Fscanf(reader, f, a...)
}

func main() {
	var n, q int
	scanf("%d\n", &n)

	mapper := make(map[string]bool)

	for i := 0; i < n; i++ {
		val, _ := reader.ReadString('\n')
		mapper[val] = false
	}

	scanf("%d\n", &q)

	var ans int
	sequneces := make(map[string]bool)

	for i := 0; i < q; i++ {
		val, _ := reader.ReadString('\n')
		sequneces[val] = true
		if len(sequneces) == n {
			ans++
			sequneces = make(map[string]bool)
			sequneces[val] = true
		}
	}

	fmt.Println(ans)
}
