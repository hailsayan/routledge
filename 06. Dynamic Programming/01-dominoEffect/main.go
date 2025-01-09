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

var dominoRes [10000]int

func main() {
	dominoRes[0] = 0
	dominoRes[1] = 1
	dominoRes[2] = 2
	var n int

	scanf("%d\n", &n)
	res := Domino(n)
	printf("%v\n", res)
	writer.Flush()
}

func Domino(n int) int {
	for i := 3; i <= n; i++ {
		dominoRes[i] = (dominoRes[i-1] + dominoRes[i-2]) % 1000000007
	}
	return dominoRes[n]
}
