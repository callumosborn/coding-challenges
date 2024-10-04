package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

// Given the head of a sorted linked list,delete all duplicates such that
// each element appears only once. Return the linked list sorted as well.
//
// Time Complexity: O(N)
// Space Complexity: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow := head
	fast := head.Next

	for fast != nil {

		if slow.Val == fast.Val {
			fast = fast.Next
			slow.Next = fast
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}

	return head
}
