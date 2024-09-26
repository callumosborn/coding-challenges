package main

import "testing"

func TestFindMaxAverage(t *testing.T) {
	_ = findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)

	_ = findMaxAverage([]int{5}, 1)
}
