package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

type m struct {
	l int
	r int
	b bool
}

func main() {
	var n int
	scanf("%d\n", &n)
	ms := make([]m, n)
	for i := 0; i < n; i++ {
		mi := new(m)
		scanf("%d %d ", &mi.l, &mi.r)
		ms[i] = *mi
	}
	sort.Slice(ms, func(i int, j int) bool {
		return ms[i].r < ms[j].r
	})
	last := ms[0].r
	res := 1
	for _, v := range ms {
		if last <= v.l {
			last = v.r
			res++
		}
	}
	printf("%d\n", res)
	defer writer.Flush()
}
