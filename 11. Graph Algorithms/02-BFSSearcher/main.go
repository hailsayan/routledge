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

const MOD int = 1e9 + 7

var g [][]int
var q []int
var distance []int
var path []int
var mark []bool

func main() {
	var n, m, s, t, v, u int
	scanf("%d %d\n", &n, &m)
	scanf("%d %d\n", &s, &t)

	mark = make([]bool, n+1)
	path = make([]int, n+1)
	g = make([][]int, n+1)
	q = make([]int, 0)
	distance = make([]int, n+1)

	for i := 1; i <= m; i++ {
		scanf("%d %d\n", &v, &u)
		g[v] = append(g[v], u)
		g[u] = append(g[u], v)
	}

	bfs(s, t)

	if !mark[t] {
		printf("%d", -1)

	} else {
		printf("%d\n", distance[t])
		ans := make([]int, 0)
		for v := t; v != -1; v = path[v] {
			ans = append(ans, v)
		}
		for i := len(ans) - 1; i >= 0; i-- {
			printf("%d ", ans[i])
		}
	}
	writer.Flush()
}

func bfs(s, t int) {
	q = append(q, s)
	mark[s] = true
	path[s] = -1
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for _, u := range g[node] {
			if !mark[u] {
				mark[u] = true
				q = append(q, u)
				distance[u] = distance[node] + 1
				path[u] = node
			}

		}
	}

}
