package main

// https://leetcode.com/problems/middle-of-the-linked-list

type ListNode struct {
	Val  int
	Next *ListNode
}

// Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.
//
// Time Complexity: O(N)
// Space Complexity: O(1)
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
