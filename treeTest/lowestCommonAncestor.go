package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	var res *TreeNode
	var traversal func(root, p, q *TreeNode) int
	traversal = func(root, p, q *TreeNode) int {
		if root == nil {
			return 0
		}
		if root == p || root == q {
			if 1+traversal(root.Left, p, q)+traversal(root.Right, p, q) == 2 {
				res = root
				return 0
			} else {
				return 1
			}
		} else {
			if traversal(root.Left, p, q)+traversal(root.Right, p, q) == 2 {
				res = root
				return 0
			}
			return traversal(root.Left, p, q) + traversal(root.Right, p, q)
		}
		return 0
	}
	traversal(root, p, q)
	return res
}
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val > p.Val && root.Val > p.Val {
		node := lowestCommonAncestor3(root.Left, p, q)
		if node != nil {
			return node
		}

	}
	if root.Val < p.Val && root.Val < p.Val {
		node := lowestCommonAncestor3(root.Right, p, q)
		if node != nil {
			return node
		}
	}
	return root
}
func lowestCommonAncestor4(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	tmpRes := 0
	var res *TreeNode
	var traversal func(root, p, q *TreeNode) int
	traversal = func(root, p, q *TreeNode) int {
		if root == nil {
			return 0
		}
		tmpRes = traversal(root.Left, p, q) + traversal(root.Right, p, q)
		if root == p || root == q {
			if 1+tmpRes == 2 {
				res = root
				return 0
			} else {
				return 1
			}
		} else {

			if tmpRes == 2 {
				res = root
				return 0
			}
			return tmpRes
		}
		return 0
	}
	traversal(root, p, q)
	return res
}
