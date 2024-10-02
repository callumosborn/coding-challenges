package main

// https://leetcode.com/problems/minimum-common-value

// Given two integer arrays nums1 and nums2, sorted in non-decreasing order,
// return the minimum integer common to both arrays.
// If there is no common integer amongst nums1 and nums2, return -1.
//
// Time Complexity: O(N + M)
// Space Complexity: O(1)
func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {

		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return -1
}
