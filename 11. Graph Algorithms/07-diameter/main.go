package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	graph := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscanf(reader, "%d %d\n", &u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	bfs := func(start int) (int, int) {
		queue := []int{start}
		dist := make([]int, n+1)
		visited := make([]bool, n+1)
		visited[start] = true

		var farthestNode, maxDist int
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					dist[neighbor] = dist[node] + 1
					queue = append(queue, neighbor)
					if dist[neighbor] > maxDist {
						maxDist = dist[neighbor]
						farthestNode = neighbor
					}
				}
			}
		}
		return farthestNode, maxDist
	}

	farthestNode1, _ := bfs(1)

	_, diameter := bfs(farthestNode1)

	fmt.Println(diameter)
}
