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
	var q, k, tmp, total int
	scanf("%d\n", &q)

	for i := 0; i < q; i++ {
		scanf("%d ", &k)
		for j := 0; j < k; j++ {
			scanf("%d ", &tmp)
			heap.Push(h, tmp+i)
		}
		l := h.Len()
		for j := 0; j < l; j++ {
			min := heap.Pop(h).(int)
			if min > i {
				total++
				break
			}
		}
	}
	printf("%d\n", total)
	writer.Flush()
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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
