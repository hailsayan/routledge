package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func buildSegmentTree(arr []int, tree []int, node int, start int, end int) {
	if start == end {
		tree[node] = arr[start]
	} else {
		mid := (start + end) / 2
		buildSegmentTree(arr, tree, 2*node+1, start, mid)
		buildSegmentTree(arr, tree, 2*node+2, mid+1, end)
		tree[node] = min(tree[2*node+1], tree[2*node+2])
	}
}

func query(tree []int, node int, start int, end int, l int, r int) int {
	if r < start || end < l {
		return math.MaxInt
	}
	if l <= start && end <= r {
		return tree[node]
	}
	mid := (start + end) / 2
	leftMin := query(tree, 2*node+1, start, mid, l, r)
	rightMin := query(tree, 2*node+2, mid+1, end, l, r)
	return min(leftMin, rightMin)
}

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d ", &arr[i])
	}
	scanf("%d\n")

	tree := make([]int, 4*n)
	buildSegmentTree(arr, tree, 0, 0, n-1)

	for i := 0; i < q; i++ {
		var l, r int
		scanf("%d %d\n", &l, &r)
		fmt.Println(query(tree, 0, 0, n-1, l, r))
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
