package general

import binarytree "github.com/freddiehaddad/leetcode/binary_tree"

type TreeNode = binarytree.TreeNode

// 106. Construct Binary Tree from Inorder and Postorder Traversal
//
// Given two integer arrays inorder and postorder where inorder is the inorder
// traversal of a binary tree and postorder is the postorder traversal of the
// same tree, construct and return the binary tree.
func buildTreeIP(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	last := len(postorder) - 1
	val := postorder[last]
	root := &TreeNode{Val: val, Left: nil, Right: nil}

	var i int
	for ; inorder[i] != val; i++ {
	}

	root.Left = buildTreeIP(inorder[:i], postorder[:i])
	root.Right = buildTreeIP(inorder[i+1:], postorder[i:last])
	return root
}

// Given two integer arrays preorder and inorder where preorder is the preorder
// traversal of a binary tree and inorder is the inorder traversal of the same
// tree, construct and return the binary tree.
func buildTreePI(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]
	root := &TreeNode{Val: val, Left: nil, Right: nil}

	var i int
	for ; inorder[i] != val; i++ {
	}

	root.Left = buildTreePI(preorder[1:i+1], inorder[:i])
	root.Right = buildTreePI(preorder[i+1:], inorder[i+1:])
	return root
}
