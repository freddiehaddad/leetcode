package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 21. Merge Two Sorted Lists
//
// You are given the heads of two sorted linked lists list1 and list2.
//
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
// Return the head of the merged linked list.
//
// Constraints:
//
//  1. The number of nodes in both lists is in the range [0, 50].
//  2. -100 <= Node.val <= 100
//  3. Both list1 and list2 are sorted in non-decreasing order.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	iter := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			iter.Next = list1
			list1 = list1.Next
		} else {
			iter.Next = list2
			list2 = list2.Next
		}
		iter = iter.Next
	}

	if list1 != nil {
		iter.Next = list1
	}

	if list2 != nil {
		iter.Next = list2
	}

	return head.Next
}

// 92. Reverse Linked List II
//
// Given the head of a singly linked list and two integers left and right where
// left <= right, reverse the nodes of the list from position left to position
// right, and return the reversed list.
//
// Example 1:
//
// Input: head = [1,2,3,4,5], left = 2, right = 4 Output: [1,4,3,2,5]
//
// Example 2:
//
// Input: head = [5], left = 1, right = 1 Output: [5]
//
// Constraints:
//
//  1. The number of nodes in the list is n.
//  2. 1 <= n <= 500
//  3. -500 <= Node.val <= 500
//  4. 1 <= left <= right <= n
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	var beforeLeft *ListNode
	reverseStart := dummy
	for ; left > 0; left, right = left-1, right-1 {
		beforeLeft = reverseStart
		reverseStart = reverseStart.Next
	}

	var previous *ListNode
	for ; right >= 0; right-- {
		temp := reverseStart.Next
		reverseStart.Next = previous
		previous = reverseStart
		reverseStart = temp
	}

	temp := beforeLeft.Next
	beforeLeft.Next = previous
	temp.Next = reverseStart
	return dummy.Next
}
