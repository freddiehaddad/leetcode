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

// 236. Lowest Common Ancestor of a Binary Tree
//
// Given a binary tree, find the lowest common ancestor (LCA) of two given
// nodes in the tree.
//
// According to the definition of LCA on Wikipedia: “The lowest common ancestor
// is defined between two nodes p and q as the lowest node in T that has both p
// and q as descendants (where we allow a node to be a descendant of itself).”
//
// Constraints:
//
//  1. The number of nodes in the tree is in the range [2, 105].
//  2. -109 <= Node.val <= 109
//  3. All Node.val are unique.
//  4. p != q
//  5. p and q will exist in the tree.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}

	var left, right *TreeNode
	if root != nil {
		left = lowestCommonAncestor(root.Left, p, q)
	}

	if root != nil {
		right = lowestCommonAncestor(root.Right, p, q)
	}

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}

// 129. Sum Root to Leaf Numbers
//
// You are given the root of a binary tree containing digits from 0 to 9 only.
// Each root-to-leaf path in the tree represents a number. For example, the
// root-to-leaf path 1 -> 2 -> 3 represents the number 123.
//
// Return the total sum of all root-to-leaf numbers. Test cases are generated
// so that the answer will fit in a 32-bit integer. A leaf node is a node with
// no children.
func sumNumbers(root *TreeNode) int {
	var sum int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left == nil && node.Right == nil {
			sum += node.Val
			continue
		}

		if node.Left != nil {
			left := node.Left
			left.Val = node.Val*10 + left.Val
			queue = append(queue, left)
		}

		if node.Right != nil {
			right := node.Right
			right.Val = node.Val*10 + right.Val
			queue = append(queue, right)
		}
	}
	return sum
}
