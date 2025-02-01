package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func power(x float64, y int) float64 {
	if int(y) == 0 {
		return 1
	} else if int(y)%2 == 0 {
		return power(x, y/2) * power(x, y/2)
	} else {
		return x * power(x, y/2) * power(x, y/2)
	}
}

func main() {
	var base float64
	var exp int
	scanf("%f\n", &base)
	scanf("%d\n", &exp)
	fmt.Printf("%.3f", power(base, exp))
}
