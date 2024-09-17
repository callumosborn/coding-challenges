package main

// https://leetcode.com/problems/two-sum

// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.
//
// You can return the answer in any order.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func twoSums(nums []int, target int) []int {
	memo := make(map[int]int)

	for i, v := range nums {

		if mi, ok := memo[target-v]; ok {
			return []int{mi, i}
		} else {
			memo[v] = i
		}
	}

	return []int{}
}
