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

type queue struct {
	array []int
	front int
	back  int
}

func newQueue(size int) queue {
	return queue{
		array: make([]int, size),
		front: 0,
		back:  0,
	}
}
func (q *queue) size() int {
	return q.back - q.front
}
func (q *queue) pushBack(x int) {
	q.array[q.back] = x
	q.back++
}
func (q *queue) getFront() int {
	return q.array[q.front]
}
func (q *queue) popFront() {
	q.front++
}

func main() {

	var n, tmp int
	scanf("%d\n", &n)
	a := make([]int, n)
	counter := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d ", &tmp)

		a[i] = tmp - 1
		counter[a[i]]++
	}
	q := newQueue(50)
	for {
		for i, v := range counter {
			// fmt.Println(i+1, v)
			if v == 0 {
				q.pushBack(i)
				// fmt.Println("*", i+1)
			}
		}

		if q.size() == 0 {
			break
		} else {
			counter[a[q.getFront()]]--
			counter[q.getFront()]--
			q.popFront()
		}
	}
	sum := 0
	for k, v := range counter {
		if v > 0 {
			sum++
			printf("%d ", k+1)
		}
	}
	fmt.Println(sum)
	writer.Flush()

}
