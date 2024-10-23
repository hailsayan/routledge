package main

import "fmt"

func main() {
	var i, j, k, n, count int64
	fmt.Scanf("%d\n", &n, &k)
	if n == 3 {
		fmt.Println(1)
		return
	}
	for i = 0; i < n/2; i++ {
		for j = i; j < n-i; j++ {
			k = n - i - j
			if i+j > k && i+k > j && j+k > i && k >= j && i+j+k == n {
				count++
			}
		}
	}
	fmt.Println(count)
}
