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
