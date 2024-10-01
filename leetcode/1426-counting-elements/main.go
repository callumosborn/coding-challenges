package main

// https://leetcode.com/problems/counting-elements/

// Given an integer array arr, count how many elements x there are,
// such that x + 1 is also in arr. If there are duplicates in arr, count them separately.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func countElements(arr []int) int {
	var count int

	hash := make(map[int]bool)

	for _, v := range arr {
		hash[v] = true
	}

	for _, v := range arr {
		if hash[v+1] {
			count++
		}
	}

	return count
}
