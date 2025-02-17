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

	var n, tmp, day int
	var name string
	scanf("%d ", &n)
	del := make([]*IntHeap, n)
	for i := 0; i < n; i++ {
		del[i] = &IntHeap{}
	}

	for i := 0; i < n; i++ {
		scanf("%d ", &tmp)
		for j := 0; j < tmp; j++ {
			scanf("%s %d ", &name, &day)
			s := sub{name, day + i - 1}
			heap.Push(h, s)
			if day+i-1 < n {
				heap.Push(del[day+i-1], s)
			}
		}
		for h.Len() > 0 {
			extra := heap.Pop(h)
			if extra.(sub).int > i {
				for j := 0; j < del[extra.(sub).int].Len(); j++ {
					if (*del[extra.(sub).int])[j].string == extra.(sub).string {
						(*del[extra.(sub).int])[j].int = i
						heap.Push(del[i], (*del[extra.(sub).int])[j])
						// println("=> ", extra.(sub).string, extra.(sub).int)
						break
					}
				}
				break
			}
		}

		for del[i].Len() > 0 {
			if (*del[i])[0].int == i {
				printf("%s ", heap.Pop(del[i]).(sub).string)
			} else {
				heap.Pop(del[i])
			}
		}
		printf("\n")
	}

	writer.Flush()

}

type sub struct {
	string
	int
}

type IntHeap []sub

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	// if h[i].int == h[j].int {
	// 	return h[i].string < h[j].string
	// }
	return h[i].string < h[j].string
}
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(sub))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
