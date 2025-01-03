package main

import "fmt"

func findMode(root *TreeNode) []int {
	res := []int{}
	var pre *TreeNode
	modeNum := 1
	maxNum := 1
	var tranversal func(root *TreeNode)
	tranversal = func(root *TreeNode) {
		if root.Left != nil {
			tranversal(root.Left)
		}
		if pre != nil {
			if pre.Val == root.Val {
				modeNum++
			} else {
				if modeNum > maxNum {
					res = res[0:0]
					res = append(res, pre.Val)
					maxNum = modeNum
				} else if modeNum == maxNum {
					res = append(res, pre.Val)
				}
				modeNum = 1
			}
		}
		pre = root
		if root.Right != nil {
			tranversal(root.Right)
		}
	}
	tranversal(root)
	fmt.Printf("mode num is %d,max num is %d", modeNum, maxNum)
	if modeNum > maxNum {
		res = res[0:0]
		res = append(res, pre.Val)
	} else if modeNum == maxNum {
		res = append(res, pre.Val)
	}
	return res
}
