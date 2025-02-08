package main

import (
	"fmt"
	"sort"
)

func validateParentheses(s string) {
	stack := []int{}        
	pairs := [][2]int{}      
	
	for i, char := range s {
		if char == '(' {
			stack = append(stack, i+1)  
		} else if char == ')' {
			if len(stack) == 0 {
				fmt.Println(-1)  
				return
			}
			openIndex := stack[len(stack)-1] 
			stack = stack[:len(stack)-1]   
			pairs = append(pairs, [2]int{openIndex, i + 1})
		}
	}

	if len(stack) > 0 {
		fmt.Println(-1)
		return
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	for _, pair := range pairs {
		fmt.Println(pair[0], pair[1])
	}
}

func main() {
	var s string
	fmt.Scan(&s)
	validateParentheses(s)
}
