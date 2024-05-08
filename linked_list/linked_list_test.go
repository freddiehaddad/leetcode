package linkedlist

import "testing"

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1    []int
		list2    []int
		expected []int
	}{
		{
			[]int{},
			[]int{},
			[]int{},
		},
		{
			[]int{},
			[]int{0},
			[]int{0},
		},
		{
			[]int{0},
			[]int{},
			[]int{0},
		},
		{
			[]int{1, 2, 4},
			[]int{1, 3, 4},
			[]int{1, 1, 2, 3, 4, 4},
		},
	}

	for i, test := range tests {
		list1 := generateLinkedList(test.list1)
		list2 := generateLinkedList(test.list2)
		expected := generateLinkedList(test.expected)
		result := mergeTwoLists(list1, list2)

		for expected != nil && result != nil {
			if expected.Val != result.Val {
				t.Errorf("[%d] value wrong. expected=%d got=%d",
					i, expected.Val, result.Val,
				)
			}
			expected = expected.Next
			result = result.Next
		}

		if expected != nil {
			t.Errorf("[%d] result wrong. missing elments:", i)
			for expected != nil {
				t.Errorf("  expected.Val=%d", expected.Val)
				expected = expected.Next
			}
		}

		if result != nil {
			t.Errorf("[%d] result wrong. extra elments:", i)
			for result != nil {
				t.Errorf("  expected.Val=%d", result.Val)
				result = result.Next
			}
		}
	}
}

func generateLinkedList(values []int) *ListNode {
	head := &ListNode{}
	iter := head

	for _, value := range values {
		iter.Next = &ListNode{Val: value}
		iter = iter.Next
	}

	return head.Next
}
