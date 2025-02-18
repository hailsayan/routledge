package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriterSize(os.Stdout, math.MaxInt32)

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	h := &IntHeap{}
	h2 := &IntHeap2{}

	var q, tmp, mid int
	mid = math.MaxInt32
	scanf("%d\n", &q)

	for i := 0; i < q; i++ {
		scanf("%d\n", &tmp)

		if tmp > mid {
			heap.Push(h, tmp)
		} else {
			heap.Push(h2, tmp)
		}

		if h2.Len() > h.Len()+1 {
			heap.Push(h, heap.Pop(h2))
		} else if h.Len() > h2.Len() {
			heap.Push(h2, heap.Pop(h))
		}

		mid = (*h2)[0]
		printf("%d\n", mid)
	}
	writer.Flush()
}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	// if h[i].int == h[j].int {
	// 	return h[i].string < h[j].string
	// }
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntHeap2 []int

func (h IntHeap2) Len() int { return len(h) }
func (h IntHeap2) Less(i, j int) bool {
	// if h[i].int == h[j].int {
	// 	return h[i].string < h[j].string
	// }
	return h[j] < h[i]
}
func (h IntHeap2) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap2) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
