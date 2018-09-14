package main

import (
	"testing"
	"fmt"
)

func TestMultiple(t *testing.T) {
	var n int = 10
	var m []int
//	var c []int
        var c = make([]int, 5)
	c[0] = 0
	c[1] = 3
	c[2] = 5
	c[3] = 6
	c[4] = 9
	m = multiple(n)
	if len(m) != len(c) {
		t.Error("Expected length:",len(c), "got length:", len(m))
//		return 
	}
	for i, v:= range m {
		if v != c[i] {
		fmt.Println(v, c[i])
		t.Error("Expected:", c[i], "got:", v)
		}
	}
}

func TestSum(t *testing.T) {
	var n int = 10
	var x = sum_m(n)
	if (x != 23) {
		t.Error("Expected 23, got:", x)
	}
}
