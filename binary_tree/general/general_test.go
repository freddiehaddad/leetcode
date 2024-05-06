package general

import (
	"slices"
	"testing"
)

func TestBuildTreeIP(t *testing.T) {
	tests := []struct {
		inorder   []int
		postorder []int
	}{
		{
			[]int{9, 3, 15, 20, 7},
			[]int{9, 15, 7, 20, 3},
		},
	}

	for i, test := range tests {
		tree := buildTreeIP(test.inorder, test.postorder)
		inorder := getInorderSlice(tree)
		if slices.Compare(test.inorder, inorder) != 0 {
			t.Errorf(
				"[%d] inorder wrong. expected=%v got=%v",
				i, test.inorder, inorder,
			)
		}
		postorder := getPostorderSlice(tree)
		if slices.Compare(test.postorder, postorder) != 0 {
			t.Errorf(
				"[%d] preorder wrong. expected=%v got=%v",
				i, test.postorder, postorder,
			)
		}

	}
}

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

func TestConnect(t *testing.T) {
	tests := []struct {
		preorder []int
		inorder  []int
		expected [][]int
	}{
		{
			[]int{1, 2, 4, 5, 3, 7},
			[]int{4, 2, 5, 1, 3, 7},
			[][]int{{1}, {2, 3}, {4, 5, 7}},
		},
	}

	for i, test := range tests {
		tree := buildTreePI(test.preorder, test.inorder)
		tree = connect(tree)
		checkBreathFirst(t, i, tree, test.expected)
	}
}

func checkBreathFirst(
	t *testing.T, ti int, root *TreeNode, expected [][]int,
) {
	var level int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		if len(queue) != len(expected[level]) {
			t.Errorf(
				"[%d] length error. level=%d expected=%d got=%d",
				ti, level, len(expected[level]), len(queue),
			)
		}

		iter := queue[0]
		for i := 0; i < len(expected[level]); i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if iter.Val != expected[level][i] {
				t.Errorf("[%d] value wrong. level=%d expected=%d got=%d",
					ti, level, expected[level][i], iter.Val,
				)
			}

			iter = iter.Next
		}
		if iter != nil {
			t.Errorf("[%d] iter wrong. level=%d expected=%v got=%v",
				ti, level, nil, iter,
			)
		}
		queue = queue[len(expected[level]):]
		level++
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

func getPostorderSlice(root *TreeNode) []int {
	var order []int
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		postorder(node.Left)
		postorder(node.Right)
		order = append(order, node.Val)
	}

	postorder(root)
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
