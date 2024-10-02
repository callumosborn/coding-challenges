package main

import "testing"

func TestGetCommon(t *testing.T) {
	_ = getCommon([]int{1, 2, 3}, []int{2, 4})

	_ = getCommon([]int{1, 2, 3}, []int{4, 5, 6})
}
