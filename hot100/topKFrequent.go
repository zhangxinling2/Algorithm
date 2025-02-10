package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]] += 1
		} else {
			m[nums[i]] = 1
		}
	}
	hp := &TopKHeap{}
	heap.Init(hp)
	for key, v := range m {
		heap.Push(hp, []int{key, v})
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(hp).([]int)[0])
	}
	return res
}

type TopKHeap [][]int

func (h TopKHeap) Len() int {
	return len(h)
}
func (h TopKHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h TopKHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *TopKHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *TopKHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 347. Top K Frequent Elements
// 347. 前 K 个高频元素
// 思路：哈希表 + 桶排序
// time O(N) space O(N)
func topKFrequent2(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	ht := make(map[int]int, 0)
	for _, v := range nums {
		ht[v]++
	}
	// 使用「桶排序」来进行频次筛选
	buckets := make([][]int, len(nums)+1)
	for num, cnt := range ht {
		if len(buckets[cnt]) == 0 {
			buckets[cnt] = make([]int, 0)
		}
		buckets[cnt] = append(buckets[cnt], num)
	}

	ans := make([]int, 0)
	for i := len(buckets) - 1; i >= 0; i-- {
		// 空桶，跳过
		if len(buckets[i]) == 0 {
			continue
		}
		ans = append(ans, buckets[i]...)
		// 已经获得 top k，则停止筛选。
		if len(ans) == k {
			break
		}
	}

	return ans
}
