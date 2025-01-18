package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var dpr [][]float64
var dpl [][]float64

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	var n, m int
	scanf("%d %d\n", &n, &m)
	arr := make([][]int, n)
	dpr = make([][]float64, n)
	dpl := make([][]float64, n)
	for i := range dpr {
		arr[i] = make([]int, m)
		dpr[i] = make([]float64, m)
		dpl[i] = make([]float64, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			scanf("%d", &arr[i][j])
		}
		scanf("\n")
	}
	dpr[0][0] = float64(arr[0][0])
	dpl[0][0] = float64(arr[0][0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var d, r float64
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				dpr[i][j] = dpr[i][j-1] + float64(arr[i][j])
				continue
			}
			d = math.Min(dpr[i-1][j], dpl[i-1][j])
			if j == 0 {
				dpr[i][j] = d + float64(arr[i][j])
				continue
			}
			r = dpr[i][j-1]
			dpr[i][j] = math.Min(float64(r), float64(d)) + float64(arr[i][j])
		}
		for j := m - 1; j >= 0; j-- {
			var d, l float64
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 && j != 0 {
				dpl[i][j] = math.Inf(1)
				continue
			}
			d = math.Min(float64(dpr[i-1][j]), float64(dpl[i-1][j]))
			if j == m-1 {
				dpl[i][j] = d + float64(arr[i][j])
				continue
			}
			l = dpl[i][j+1]
			dpl[i][j] = math.Min(float64(l), float64(d)) + float64(arr[i][j])
		}
	}
	fmt.Println(int64(dpr[n-1][m-1]))
	defer writer.Flush()
}
