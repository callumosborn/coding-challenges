package main

// https://leetcode.com/problems/jewels-and-stones/

// You're given strings jewels representing the types of stones that are jewels,
// and stones representing the stones you have.
// Each character in stones is a type of stone you have.
// You want to know how many of the stones you have are also jewels.
//
// Time Complexity: O(N + M)
// Space Complexity: O(N)
func numJewelsInStones(jewels string, stones string) int {
	counts := make(map[rune]bool)

	for _, v := range jewels {
		counts[v] = true
	}

	var ans int

	for _, v := range stones {

		if counts[v] {
			ans++
		}
	}

	return ans
}
