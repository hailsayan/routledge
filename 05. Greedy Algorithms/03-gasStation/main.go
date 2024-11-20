package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstLine, _ := reader.ReadString('\n')
	parts := strings.Fields(firstLine)
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])
	L, _ := strconv.Atoi(parts[2])

	secondLine, _ := reader.ReadString('\n')
	parts = strings.Fields(secondLine)

	stations := make([]int, n+1) 
	for i := 0; i < n; i++ {
		stations[i], _ = strconv.Atoi(parts[i])
	}
	stations[n] = L 

	sort.Ints(stations)

	currentFuel := k
	lastRefillIndex := -1
	refills := []int{}

	for i := 0; i < len(stations); {
		maxReach := currentFuel
		bestIndex := -1

		for j := i; j < len(stations) && stations[j] <= maxReach; j++ {
			bestIndex = j
		}

		if bestIndex == -1 || bestIndex == lastRefillIndex {
			fmt.Println(-1)
			return
		}

		if bestIndex != n { 
			refills = append(refills, bestIndex+1) 
		}

		currentFuel = stations[bestIndex] + k
		lastRefillIndex = bestIndex
		i = bestIndex + 1
	}

	fmt.Println(len(refills))
	if len(refills) > 0 {
		fmt.Println(strings.Trim(fmt.Sprint(refills), "[]"))
	}
}