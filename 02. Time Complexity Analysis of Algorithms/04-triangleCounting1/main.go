package main

import "fmt"

func main() {
	var i, j, k, n, count int64
	fmt.Scanf("%d\n", &n, &k)
	for i = 0; i < n; i++ {
		for j = i; j < n; j++ {
			for k = j; k < n; k++ {
				if i+j > k && i+k > j && j+k > i && i+j+k == n {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
