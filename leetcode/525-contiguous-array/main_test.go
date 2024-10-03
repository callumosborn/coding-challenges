package main

import "testing"

func TestFindMaxLength(t *testing.T) {
	_ = findMaxLength([]int{0, 1})

	_ = findMaxLength([]int{0, 1, 0})

	_ = findMaxLength([]int{0, 0, 1, 0, 1, 0, 0, 1})
}
