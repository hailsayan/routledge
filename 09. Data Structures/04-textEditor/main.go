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
	array   []int
	front   int
	back    int
	maxSize int
}

func newQueue(size int) queue {
	return queue{
		array:   make([]int, size),
		front:   0,
		back:    0,
		maxSize: size,
	}
}
func (q *queue) size() int {
	return q.back - q.front
}

func (q *queue) pushBack(x int) {
	q.array[q.back] = x
	q.back = (q.back + 1) % q.maxSize
}
func (q *queue) pushFront(x int) {
	q.front = (q.front - 1 + q.maxSize) % q.maxSize
	q.array[q.front] = x
}

func (q *queue) getFront() int {
	return q.array[q.front]
}
func (q *queue) getBack() int {
	q.back = (q.back - 1 + q.maxSize) % q.maxSize
	return q.array[q.back-1]
}

func (q *queue) popFront() {
	q.front++
}
func (q *queue) popBack() {
	q.back--
}

func main() {

	var n int
	var tmp string
	var ch rune
	scanf("%d", &n)
	a := make([]rune, 0)
	var pointer int

	for i := 0; i < n; i++ {
		scanf("\n%s", &tmp)
		if tmp == "+" {
			if pointer < len(a) {
				pointer++
			}
		} else if tmp == "-" {
			if pointer > 0 {
				pointer--
			}
		} else {
			scanf(" %c", &ch)
			// a = insert(a, pointer, tmp2)
			if pointer == len(a) {
				a = append(a, ch)
			} else {
				for j := pointer; j < len(a); j++ {
					ch, a[j] = a[j], ch
				}
				a = append(a, ch)
			}

			pointer++
		}
	}

	for _, c := range a {
		printf("%c", c)
	}
	writer.Flush()

}

func insert(a []rune, i int, x rune) []rune {
	if len(a) == 0 {
		return []rune{x}
	}
	new := make([]rune, len(a)+1)
	for j := 0; j < len(a)+1; j++ {
		if j < i {
			new[j] = a[j]
		} else if j == i {
			new[j] = x
		} else {
			new[j] = a[j-1]
		}
	}

	return new

}
