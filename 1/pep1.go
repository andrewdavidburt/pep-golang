package main

import (
	"fmt"
)

func multiple(n int) []int {
	var output []int
	for i := 0; i < n; i++ {
		if (i%5 == 0) || (i%3 == 0) {
//			fmt.Println("multiples:", i)
			output = append(output, i)
		}
	}
	return output
}

func sum_m(n int) int {
	var m []int
	var sum int
	m = multiple(n)
	for i := range m {
//		fmt.Println("summed:", i)
		sum += m[i]
	}
	return sum
}

func main () {
	var n int = 1000
	var x = sum_m(n)
	fmt.Println("For starting value:", n, "result is: ", x)
}
