package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func recursivePartition(idx int, sum1 int, sum2 int, nums []int) int {
	if idx == len(nums) {
		return int(math.Abs(float64(sum1 - sum2)))
	}
	include := recursivePartition(idx+1, sum1+nums[idx], sum2, nums)
	exclude := recursivePartition(idx+1, sum1, sum2+nums[idx], nums)
	return int(math.Min(float64(include), float64(exclude)))
}

func bitmaskPartition(nums []int) int {
	n := len(nums)
	total := 1 << n
	minDiff := math.MaxInt64

	for mask := 0; mask < total; mask++ {
		sum1, sum2 := 0, 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) > 0 {
				sum1 += nums[i]
			} else {
				sum2 += nums[i]
			}
		}
		minDiff = int(math.Min(float64(minDiff), math.Abs(float64(sum1-sum2))))
	}
	return minDiff
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	firstLine, _ := reader.ReadString('\n')
	firstLine = strings.TrimSpace(firstLine)
	n, _ := strconv.Atoi(firstLine)

	secondLine, _ := reader.ReadString('\n')
	secondLine = strings.TrimSpace(secondLine)
	parts := strings.Split(secondLine, " ")

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(parts[i])
	}

	// fmt.Println(recursivePartition(0, 0, 0, nums))
	fmt.Println(bitmaskPartition(nums))
}
