package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}
func inorder(root *TreeNode) (res []int) {
	cur := root
	st := list.New()
	for cur != nil || st.Len() > 0 {
		if cur != nil {
			st.PushBack(cur)
			cur = cur.Left
		} else {
			cur = st.Remove(st.Back()).(*TreeNode)
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}
func postorder(root *TreeNode) (res []int) {
	rootStack := []*TreeNode{}
	rootStack = append(rootStack, root)
	st := list.New()
	st.PushBack(root)
	for len(rootStack) != 0 {
		stackLen := len(rootStack)
		tmpNode := rootStack[stackLen-1]
		rootStack = rootStack[:stackLen-1]
		res = append(res, tmpNode.Val)
		if tmpNode.Right != nil {
			rootStack = append(rootStack, tmpNode.Left)
		}
		if tmpNode.Left != nil {
			rootStack = append(rootStack, tmpNode.Right)
		}
	}
	reverse(res)
	return res
}
func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l, r = l+1, r-1
	}
}
func preorder(root *TreeNode) (res []int) {
	rootStack := []*TreeNode{}
	rootStack = append(rootStack, root)
	st := list.New()
	st.PushBack(root)
	for len(rootStack) != 0 {
		stackLen := len(rootStack)
		tmpNode := rootStack[stackLen-1]
		rootStack = rootStack[:stackLen-1]
		res = append(res, tmpNode.Val)
		if tmpNode.Right != nil {
			rootStack = append(rootStack, tmpNode.Right)
		}
		if tmpNode.Left != nil {
			rootStack = append(rootStack, tmpNode.Left)
		}
	}
	return res
}
func levelOrder(root *TreeNode) (res []int) {
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		res = append(res, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return res
}
func levelOrder2(root *TreeNode) [][]int {
	arr := [][]int{}
	depth := 0
	var levelorder func(root *TreeNode, depth int)
	levelorder = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(arr) == depth {
			arr = append(arr, []int{})
		}
		arr[depth] = append(arr[depth], root.Val)
		levelorder(root.Left, depth+1)
		levelorder(root.Right, depth+1)
	}
	levelorder(root, depth)
	return arr
}
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	arr := [][]int{}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		tmpArr := []int{}
		l := queue.Len()
		for i := 0; i < l; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			tmpArr = append(tmpArr, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		arr = append(arr, tmpArr)
	}
	return arr
}
func levelOrder4(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	curLevel := []*TreeNode{root}
	res := [][]int{}
	for len(curLevel) > 0 {
		nextLevel := []*TreeNode{}
		tmpRes := []int{}
		for _, v := range curLevel {
			tmpRes = append(tmpRes, v.Val)
			if v.Left != nil {
				nextLevel = append(nextLevel, v.Left)
			}
			if v.Right != nil {
				nextLevel = append(nextLevel, v.Right)
			}
		}
		res = append(res, tmpRes)
		curLevel = nextLevel
	}
	return res
}
func levelOrderII(root *TreeNode) [][]int {
	res := levelOrder3(root)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
func preorderII(root *TreeNode) (res []int) {
	stack := list.New()
	if root != nil {
		stack.PushBack(root)
	}
	var node *TreeNode
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		if e.Value == nil {
			node = stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, node.Val)
			continue
		}
		node = e.Value.(*TreeNode)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		stack.PushBack(node)
		stack.PushBack(nil)
	}
	return res
}
func main() {
	p := &TreeNode{Val: 7}
	q := &TreeNode{
		Val: 4,
	}
	t1 := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
			Right: q,
		},
		Right: &TreeNode{Val: 2,
			Right: &TreeNode{
				Val:  5,
				Left: p,
			}},
	}

	//fmt.Println(levelOrderII(t1))
	//二叉树的右视图
	//fmt.Println(rightSideView(t1))
	//翻转二叉树
	//invertTreeII(t1)
	//fmt.Println(preorder(t1))
	//对称二叉树
	//fmt.Println(isSymmetric(t1))
	//二叉树的最大深度
	//fmt.Println(maxDepth(t1))
	//二叉树的最小深度
	//fmt.Println(minDepth(t1))
	//是否是平衡二叉树
	//fmt.Println(isBalanced(t1))
	//二叉树的所有路径
	//fmt.Println(allPath(t1))
	//左叶子节点之和
	//fmt.Println(leftLeafSum(t1))
	//路径总和
	//fmt.Println(hasPath(t1, 13))
	//最大二叉树
	//fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
	//是否是二叉搜索树
	//fmt.Println(isBinarySearchTree(t2))
	//t2 := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 0,
	//		Left: &TreeNode{
	//			Val: 0,
	//		},
	//		Right: &TreeNode{
	//			Val: 0,
	//		},
	//	},
	//	Right: &TreeNode{Val: 1,
	//		Left: &TreeNode{Val: 1},
	//		Right: &TreeNode{
	//			Val: 1,
	//		}},
	//}
	//找二叉搜索树的额最大众数
	//fmt.Println(findMode(t2))
	//找二叉树最小公共节点
	fmt.Println(lowestCommonAncestor2(t1, p, q))
	//有序数组构建二叉搜索树
	fmt.Println(buildBinaryTreeFromNums([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
