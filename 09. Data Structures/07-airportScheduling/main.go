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
	var q, k, tmp int
	scanf("%d %d\n", &q, &k)
	a := make([]int, 0)
	for i := 0; i < q; i++ {
		scanf("%d\n", &tmp)
		b := sort.Search(len(a), func(x int) bool { return a[x] > tmp })
		if (b == len(a) || a[b]-tmp >= k) && (b == 0 || tmp-a[b-1] >= k) {
			printf("Permission Granted!\n")
			a = append(a, tmp)
			sort.Ints(a)
		} else {
			printf("Permission Denied!\n")
		}

	}
	writer.Flush()
}