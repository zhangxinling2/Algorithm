package main

import (
	"container/list"
	"sort"
)

func queueReconstructionByHeight(people [][]int) [][]int {
	// 先将身高从大到小排序，确定最大个子的相对位置
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] // 当身高相同时，将K按照从小到大排序
		}
		return people[i][0] > people[j][0] // 身高按照由大到小的顺序来排
	})
	// 再按照K进行插入排序，优先插入K小的
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1]) // 空出一个位置
		people[p[1]] = p
	}
	return people
}

// 链表实现
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] //当身高相同时，将K按照从小到大排序
		}
		return people[i][0] > people[j][0]
	})
	l := list.New() //创建链表
	for i := 0; i < len(people); i++ {
		position := people[i][1]
		mark := l.PushBack(people[i]) //插入元素
		e := l.Front()
		for position != 0 { //获取相对位置
			position--
			e = e.Next()
		}
		l.MoveBefore(mark, e) //移动位置

	}
	res := [][]int{}
	for e := l.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.([]int))
	}
	return res
}
