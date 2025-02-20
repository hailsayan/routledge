package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var writer *bufio.Writer = bufio.NewWriterSize(os.Stdout, math.MaxInt32)

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	var w, h, q, tmp int
	var str string
	scanf("%d %d %d\n", &w, &h, &q)

	v := make([]int, 0, q)
	vm := make(map[int]int, q)
	v = append(v, 0)
	vm[0] = w

	r := make([]int, 0, q)
	rm := make(map[int]int, q)
	r = append(r, 0)
	rm[0] = h

	tool := int64(w)
	arz := int64(h)

	for i := 0; i < q; i++ {
		scanf("%s %d\n", &str, &tmp)
		if str == "V" {
			first := sort.SearchInts(v, tmp)
			first--
			vm[tmp] = v[first] + vm[v[first]] - tmp
			cond := vm[v[first]] == int(tool)
			vm[v[first]] -= vm[tmp]
			v = append(v, tmp)
			mySort(v, first+1)
			if cond {
				tool = bigGap(vm)
			}
		} else {
			first := sort.SearchInts(r, tmp)
			first--
			rm[tmp] = r[first] + rm[r[first]] - tmp
			cond := rm[r[first]] == int(arz)
			rm[r[first]] -= rm[tmp]
			r = append(r, tmp)
			mySort(r, first+1)
			if cond {
				arz = bigGap(rm)
			}
		}

		printf("%d\n", tool*arz)

	}
	writer.Flush()
}

func bigGap(v map[int]int) int64 {
	max := 0
	for _, v := range v {
		if v > max {
			max = v
		}
	}
	return int64(max)
}
func mySort(a []int, i int) {
	for j := i; j < len(a)-1; j++ {
		a[len(a)-1], a[j] = a[j], a[len(a)-1]
	}
}