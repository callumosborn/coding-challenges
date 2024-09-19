package main

import "testing"

func TestSortedSquares(t *testing.T) {
	_ = sortedSquares([]int{-4, -1, 0, 3, 10})

	_ = sortedSquares([]int{-7, -3, 2, 3, 11})

	_ = sortedSquares([]int{-12, -7, -3, 2, 3, 11})
}
