package main

import "testing"

func TestGetAverages(t *testing.T) {
	_ = getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3)
}
