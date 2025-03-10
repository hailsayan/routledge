package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100005

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

var (
	adj     [N][]int
	visited [N]bool
)

func dfs(u, t int, r *int) {
	visited[u] = true
	if u == t {
		fmt.Println("YES")
		*r = 1
	}
	for _, v := range adj[u] {
		if !visited[v] {
			dfs(v, t, r)
		}
	}
}

func main() {
	var n, m, s, t int
	scanf("%d %d\n", &n, &m)
	scanf("%d %d\n", &s, &t)
	for i := 0; i < m; i++ {
		var u, v int
		scanf("%d %d\n", &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	var r int
	dfs(s, t, &r)
	if r != 1 {
		fmt.Println("NO")
	}
}
