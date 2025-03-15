package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriterSize(os.Stdout, math.MaxInt32)

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

var g [][]int
var mark []bool

func main() {

	var n, x, y int
	type cell struct {
		x int
		y int
	}

	scanf("%d ", &n)
	cells := make([]cell, n)
	mark = make([]bool, n)

	for i := 0; i < n; i++ {
		scanf("%d %d ", &x, &y)
		cells[i] = cell{x: x, y: y}
		g = append(g, make([]int, 0))
	}
	var t int = 0
	for i := 0; i < n-1; i++ {
		s := t + 1
		for j := i + 1; j < n; j++ {
			if cells[i].x == cells[j].x || cells[i].y == cells[j].y {
				g[t] = append(g[t], s)
				g[s] = append(g[s], t)
			}
			s++
		}
		t++
	}
	var ans int
	for i := 0; i < len(g); i++ {
		if !mark[i] {
			dfs(i)
			ans++
		}
	}
	printf("%d", ans-1)
	writer.Flush()
}

func dfs(s int) bool {
	mark[s] = true
	for _, v := range g[s] {
		if !mark[v] && dfs(v) {
			return true
		}
	}
	return false
}
