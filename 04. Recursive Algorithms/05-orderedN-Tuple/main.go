package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var items []int
var n int

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	scanf("%d", &n)
	// start := time.Now()
	items = make([]int, n)
	for i := range items {
		items[i] = 1
	}
	printN()
	// fmt.Println(time.Since(start))
	defer writer.Flush()

}

func printN() {
	var flag = false
	for j := range items {
		if items[j] != n {
			flag = true
		}
	}
	if !flag {
		printf("%s\n", strings.Trim(fmt.Sprintf("%v", items), "[]"))
		return
	}
	printf("%s\n", strings.Trim(fmt.Sprintf("%+v", items), "[]"))
	increment(items)
	printN()
}
func increment(l []int) {
	for i := range l {
		v := l[n-i-1]
		if v < n {
			l[n-i-1]++
			break
		}
		l[n-i-1] = 1
	}
}
