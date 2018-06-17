package main

import "testing"

func TestSum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{1, 2}, 3},
		test{[]int{3, 4}, 7},
		test{[]int{5, 6, 7}, 18},
		test{[]int{-2, 2, 1}, 1},
	}

	for _, v := range tests {
		got := sum(v.data...)
		if got != v.answer {
			t.Error("expect:", v.answer, "received:", got)
		}
	}
}
