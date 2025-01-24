package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

var arr []float64
var res int
var k int

func main() {
	var n int
	scanf("%d%d", &n, &k)
	arr = make([]float64, n)
	for i := range arr {
		arr[i] = float64(i + 1)
	}
	permute(0, n)
	fmt.Println(res)
}

func permute(l, r int) {
	if l == r {
		t := 0
		for i := range arr {
			for j := i + 1; j < len(arr); j++ {
				if i < j && arr[j] < arr[i] {
					t++
				}
			}
		}
		if t == k {
			res++
		}
		return
	}
	for i := l; i < r; i++ {
		arr[l], arr[i] = arr[i], arr[l]
		permute(l+1, r)
		arr[l], arr[i] = arr[i], arr[l]
	}
}
