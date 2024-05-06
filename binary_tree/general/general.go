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

// 114. Flatten Binary Tree to Linked List
//
// Given the root of a binary tree, flatten the tree into a "linked list":
//
//  1. The "linked list" should use the same TreeNode class where the right
//     child pointer points to the next node in the list and the left child
//     pointer is always null.
//  2. The "linked list" should be in the same order as a pre-order
//     traversal of the binary tree.
func flatten(root *TreeNode) {
	for root != nil {
		if root.Left != nil {
			iter := root.Left
			for iter.Right != nil {
				iter = iter.Right
			}

			iter.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		root = root.Right
	}
}

// 112. Path Sum
//
// Given the root of a binary tree and an integer targetSum, return true if the
// tree has a root-to-leaf path such that adding up all the values along the
// path equals targetSum.
//
// A leaf is a node with no children.
//
// Constraints:
//
//  1. The number of nodes in the tree is in the range [0, 5000].
//  2. -1000 <= Node.val <= 1000
//  3. -1000 <= targetSum <= 1000
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}

	if root.Left != nil && hasPathSum(root.Left, targetSum) {
		return true
	}

	if root.Right != nil && hasPathSum(root.Right, targetSum) {
		return true
	}

	return false
}
