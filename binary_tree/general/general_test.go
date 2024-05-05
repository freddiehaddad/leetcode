package general

import (
	"slices"
	"testing"
)

func TestBuildTreePI(t *testing.T) {
	tests := []struct {
		preorder []int
		inorder  []int
	}{
		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
		},
	}

	for i, test := range tests {
		tree := buildTreePI(test.preorder, test.inorder)
		preorder := getPreorderSlice(tree)
		if slices.Compare(test.preorder, preorder) != 0 {
			t.Errorf(
				"[%d] preorder wrong. expected=%v got=%v",
				i, test.preorder, preorder,
			)
		}

		inorder := getInorderSlice(tree)
		if slices.Compare(test.inorder, inorder) != 0 {
			t.Errorf(
				"[%d] inorder wrong. expected=%v got=%v",
				i, test.inorder, inorder,
			)
		}
	}
}

func getInorderSlice(root *TreeNode) []int {
	var order []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		order = append(order, node.Val)
		inorder(node.Right)
	}

	inorder(root)
	return order
}

func getPreorderSlice(root *TreeNode) []int {
	var order []int
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		order = append(order, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}

	preorder(root)
	return order
}
