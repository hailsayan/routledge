package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

var g []*list.List
var mark []bool

var distance []float64

func main() {
	var n int
	scanf("%d\n", &n)

	mark = make([]bool, n)
	g = make([]*list.List, n)
	distance = make([]float64, n)
	for i := range distance {
		distance[i] = math.Inf(1)
	}

	distance[0] = 0

	for i := range g {
		g[i] = list.New()
	}

	for i := 0; i < n-1; i++ {
		var u, v int
		scanf("%d %d\n", &u, &v)
		g[u-1].PushBack(v - 1)
	}
	var q int
	scanf("%d\n", &q)
	friends := make([]int, q)
	for i := 0; i < q; i++ {
		var j int
		scanf("%d\n", &j)
		friends[i] = j - 1
	}
	BFS(0)
	min := 1000000
	var res int
	for i := range friends {
		if distance[friends[i]] < float64(min) {
			res = friends[i]
			min = int(distance[friends[i]])
		}
	}
	fmt.Println(res + 1)
}

var queue = make([]int, 0)

func BFS(r int) {
	queue = append(queue, r)
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for el := g[v].Front(); el != nil; el = el.Next() {
			u := el.Value.(int)
			if distance[u] > distance[v]+1 {
				distance[u] = distance[v] + 1
				queue = append(queue, u)
			}
		}
	}
}
