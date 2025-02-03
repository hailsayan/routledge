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
	var t, tmpn, tmpk, tmp int
	var arr [][]int
	var arrk []int
	scanf("%d\n", &t)
	arr = make([][]int, t)
	arrk = make([]int, t)
	for i := 0; i < t; i++ {
		scanf("%d %d\n", &tmpn, &tmpk)
		arr[i] = make([]int, tmpn+1)
		for j := 1; j <= tmpn; j++ {
			scanf("%d ", &tmp)
			arr[i][j] = arr[i][j-1] + tmp
		}
		arrk[i] = tmpk
	}

	for index, subarr := range arr {
		ans := 0
		sort.Ints(subarr)
		for x := 0; x < len(subarr); x++ {
			l := sort.SearchInts(subarr, subarr[x]-arrk[index])
			r := sort.SearchInts(subarr, subarr[x]+arrk[index]+1)
			ans += len(subarr) - (r - l)
		}

		printf("%d\n", ans/2)

	}
	writer.Flush()

}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
