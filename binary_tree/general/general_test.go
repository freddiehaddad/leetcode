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

func TestFlatten(t *testing.T) {
	tests := []struct {
		preorder []int
		inorder  []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{3, 2, 4, 1, 5, 6},
		},
	}

	for i, test := range tests {
		tree := buildTreePI(test.preorder, test.inorder)
		flatten(tree)
		for _, val := range test.preorder {
			if tree == nil {
				t.Errorf(
					"[%d] node is nil. expected=%d", i, val,
				)
				continue
			}
			if tree.Val != val {
				t.Errorf(
					"[%d] value wrong. expected=%d got=%d",
					i, val, tree.Val,
				)
			}
			if tree.Left != nil {
				t.Errorf(
					"[%d] left wrong. expected=%v got=%v",
					i, nil, tree,
				)
			}
			tree = tree.Right
		}
		if tree != nil {
			t.Errorf(
				"[%d] node not nil. expected=%v got=%v",
				i, nil, tree,
			)
			continue
		}

	}
}

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		preorder []int
		inorder  []int
		input    int
		expected bool
	}{
		{
			[]int{5, 4, 11, 7, 2, 8, 13, 4, 1},
			[]int{7, 11, 2, 4, 5, 13, 8, 4, 1},
			22,
			true,
		},
	}

	for i, test := range tests {
		tree := buildTreePI(test.preorder, test.inorder)
		result := hasPathSum(tree, test.input)
		if test.expected != result {
			t.Errorf("[%d] result wrong. expected=%t got=%t",
				i, test.expected, result,
			)
		}
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		preorder []int
		inorder  []int
		p        int
		q        int
		a        int
	}{
		{
			[]int{3, 5, 6, 2, 7, 4, 1, 0, 8},
			[]int{6, 5, 7, 2, 4, 3, 0, 1, 8},
			2,
			0,
			3,
		},
	}

	for i, test := range tests {
		tree := buildTreePI(test.preorder, test.inorder)
		p := findNode(tree, test.p)
		q := findNode(tree, test.q)
		a := findNode(tree, test.a)
		result := lowestCommonAncestor(tree, p, q)
		if a != result {
			t.Errorf("[%d] result wrong. expected=%v got=%v",
				i, a, result,
			)
		}
	}
}

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		preorder []int
		inorder  []int
		expected int
	}{
		{
			[]int{4, 9, 5, 1, 0},
			[]int{5, 9, 1, 4, 0},
			1026,
		},
	}

	for i, test := range tests {
		tree := buildTreePI(test.preorder, test.inorder)
		result := sumNumbers(tree)
		if test.expected != result {
			t.Errorf("[%d] result wrong. expected=%d got=%d",
				i, test.expected, result,
			)
		}
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

func findNode(node *TreeNode, val int) *TreeNode {
	if node.Val == val {
		return node
	}

	var left *TreeNode
	if node.Left != nil {
		left = findNode(node.Left, val)
	}

	if left != nil {
		return left
	}

	var right *TreeNode
	if node.Right != nil {
		right = findNode(node.Right, val)
	}

	return right
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
