package main

// https://leetcode.com/problems/squares-of-a-sorted-array/

// Given an integer array nums sorted in non-decreasing order,
// return an array of the squares of each number sorted in non-decreasing order.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func sortedSquares(nums []int) []int {
	aux := make([]int, len(nums))

	i, j := 0, len(nums)-1

	for i <= j {

		if nums[i]*nums[i] > nums[j]*nums[j] {
			aux[j-i] = nums[i] * nums[i]
			i++
		} else {
			aux[j-i] = nums[j] * nums[j]
			j--
		}
	}

	return aux
}
