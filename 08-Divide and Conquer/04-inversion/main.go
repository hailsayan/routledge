package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func mergeSort(arr []int, n int) int {
	if n <= 1 {
		return 0
	}
	mid := n / 2
	cnt := mergeSort(arr[:mid], mid) + mergeSort(arr[mid:], n-mid)
	i, j := 0, mid
	for i < mid && j < n {
		if arr[i] > arr[j] {
			cnt += mid - i
			j++
		} else {
			i++
		}
	}
	sort.Ints(arr)
	return cnt
}


func main() {
	var n int
	scanf("%d\n", &n)

	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d", &sequence[i])
	}
	res := mergeSort(sequence, n)
	fmt.Println(res)
}
