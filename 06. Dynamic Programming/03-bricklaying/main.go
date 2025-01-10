package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	var n int
	var set [100001]int
	set[0] = 1
	set[1] = 1
	set[2] = 1
	set[3] = 2

	m := 1000000007
	for i := 4; i < len(set); i++ {
		a := set[i-1] + set[i-2] + set[i-3] - set[i-4]
		set[i] = ((a % m) + m) % m
	}
	scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		var temp int
		scanf("%d\n", &temp)
		printf("%v\n", set[temp])
		writer.Flush()
	}
}
