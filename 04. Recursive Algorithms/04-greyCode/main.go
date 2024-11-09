package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var i = 1

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	var n int = 5

	// scanf("%d", &n)
	fmt.Print(printCode(n))
}

func printCode(n int) string {
	if n == 1 {
		return "0\n1"
	}
	s := strings.Split(printCode(n-1), "\n")
	d := make([]string, len(s))
	copy(d, s)
	for i := range s {
		d[i] = "0" + d[i]
	}
	for i := range s {
		elmnt := "1" + s[len(s)-i-1]
		d = append(d, elmnt)
	}
	return strings.Join(d, "\n")
}
