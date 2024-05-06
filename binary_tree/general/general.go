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

// 117. Populating Next Right Pointers in Each Node II
//
// Populate each next pointer to point to its next right node. If there is no
// next right node, the next pointer should be set to nil.
//
// Initially, all next pointers are set to nil.
func connect(root *TreeNode) *TreeNode {
	var c, p, s *TreeNode

	c = root
	for c != nil {
		for c != nil {
			if c.Left != nil {
				if s == nil {
					s = c.Left
				}
				if p != nil {
					p.Next = c.Left
				}
				p = c.Left
			}
			if c.Right != nil {
				if s == nil {
					s = c.Right
				}
				if p != nil {
					p.Next = c.Right
				}
				p = c.Right
			}
			c = c.Next
		}
		c = s
		p = nil
		s = nil
	}

	return root
}
