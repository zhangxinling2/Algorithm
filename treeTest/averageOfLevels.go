package main

import "container/list"

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := []float64{}
	queue := list.New()
	queue.PushBack(root)
	var sum int
	for queue.Len() > 0 {
		//保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			// 当前层元素求和
			sum += node.Val
		}
		// 计算每层的平均值，将结果添加到响应结果中
		res = append(res, float64(sum)/float64(length))
		sum = 0 // 清空该层的数据
	}
	return res
}
