package main

import (
	"testing"
	"fmt"
)

func Test_fib(t *testing.T) {
	var f []int
	var n int=10
	var l int = 100
	var c = make([]int, 10)
	c[0] = 1
	c[1] = 2
	c[2] = 3
	c[3] = 5
	c[4] = 8
	c[5] = 13
	c[6] = 21
	c[7] = 34
	c[8] = 55
	c[9] = 89
	f = fib(n, l)
	if len(c) != len(f) {
		t.Error("Expected length:", len(c), "got length:", len(f))
	}
	for i, v:= range f {
		if v != c[i] {
			fmt.Println(v, c[i])
			t.Error("Expected:", c[i], "got:", v)
		}
	}
}
		
func Test_addevens(t *testing.T) {
	var f []int
	var n int=10
	var l int = 100
	f = fib(n, l)
	var ae int
	ae = addevens(f)
	if (ae != 231) {
		t.Error("Expected 231, got:", ae)
	}
}
