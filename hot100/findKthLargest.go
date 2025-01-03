package main

import (
	"container/heap"
)

func findKthLargest(nums []int, k int) int {
	hp := &MinHeap{}
	heap.Init(hp)
	for i := 0; i < len(nums); i++ {
		heap.Push(hp, nums[i])
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}
	return heap.Pop(hp).(int)
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func findKthLargest2(nums []int, k int) int {
	quickSort(nums, k, 0, len(nums)-1)
	return nums[k-1]
}
func Partition(nums []int, start, stop int) int {
	if start >= stop {
		return -1
	}
	pivot := nums[start]
	for start < stop {
		for start < stop && nums[stop] <= pivot {
			stop--
		}
		nums[start] = nums[stop]
		for start < stop && nums[start] > pivot {
			start++
		}
		nums[stop] = nums[start]
	}
	nums[start] = pivot
	return start
}
func quickSort(nums []int, k, start, stop int) {
	if start < stop {
		index := Partition(nums, start, stop)
		if index == k-1 {
			return
		} else if index < k {
			quickSort(nums, k, index+1, stop)
		} else {

			quickSort(nums, k, start, stop-1)
		}
	}
}
