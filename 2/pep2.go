package main

import (
	"fmt"
)

func fib(n int, l int) []int {
	var output []int
	var addon int
	output = append(output, 1)
	output = append(output, 2)
	for i := 2; i < n; i++ {
		addon = output[i-2]+output[i-1]
		if (addon < l) {
			output = append(output, addon)
		} else {
			return output
		}
	}
	return output
}

func addevens(fib []int) int {
	var output int = 0
	for i := 0; i < len(fib); i++ {
		if (fib[i]%2 == 0) {
			output = output + fib[i]
		}
	}
	return output
}

func main () {
	var n int = 50
	var f []int
	var l int = 4000000
	f = fib(n, l)
	fmt.Println("Fibonacci sequence:", f)
	var ae int
	ae = addevens(f)
	fmt.Println("Fibonacci even sum:", ae)
}
