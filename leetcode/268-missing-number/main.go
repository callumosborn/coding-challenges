package main

// https://leetcode.com/problems/missing-number

// Given an array nums containing n distinct numbers in the range [0, n],
// return the only number in the range that is missing from the array.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func missingNumber(nums []int) int {
	hash := make(map[int]bool)

	for _, num := range nums {
		hash[num] = true
	}

	for i := 0; i <= len(nums); i++ {

		if !hash[i] {
			return i
		}
	}

	return -1
}
