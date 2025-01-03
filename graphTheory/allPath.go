package main

import "fmt"

type MyLink struct {
	val  int
	next *MyLink
}

func (l *MyLink) addAtHead(val int) {
	nl := &MyLink{
		val:  val,
		next: l.next,
	}
	l.next = nl
}
func (l *MyLink) addAtTail(val int) {
	t := l
	for t.next != nil {
		t = t.next
	}
	t.next = &MyLink{
		val:  val,
		next: nil,
	}
}
func (l *MyLink) get(val int) int {
	t := l
	i := 1
	for t != nil && t.next != nil {
		if t.next.val == val {
			return i
		} else {
			t = t.next
		}
	}
	return 0
}
func (l *MyLink) deleteAtIndex(val int) {
	t := l
	for i := 0; i < val; i++ {
		t = t.next
	}
	if t.next == nil {
		return
	}
	tmp := t.next.next
	t.next = nil
	t.next = tmp
}
func Construct() *MyLink {
	return &MyLink{
		val:  0,
		next: nil,
	}
}

func AllPath() {
	var (
		n int
		m int
		a int
		b int
	)
	resPath := make([][]int, 0)
	path := make([]int, 0)
	fmt.Scanln(&n, &m)
	graph := make([]*MyLink, n+1)
	for i := 0; i < n+1; i++ {
		graph[i] = &MyLink{}
	}
	for i := 0; i < m; i++ {
		_, err := fmt.Scanln(&a, &b)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		fmt.Printf("Read edge: %d -> %d\n", a, b) // 调试信息，打印读取的边
		graph[a].addAtHead(b)
	}
	var dfs func(graph []*MyLink, i int)
	dfs = func(graph []*MyLink, i int) {
		if i == n {
			tmp := make([]int, len(path))
			copy(tmp, path)
			resPath = append(resPath, tmp)
			return
		}
		node := graph[i].next
		for node != nil {
			path = append(path, node.val)
			dfs(graph, node.val)
			node = node.next
			path = path[:len(path)-1]
		}
	}
	dfs(graph, 1)
	for i := 0; i < len(resPath); i++ {
		fmt.Print("1 ")
		for j := 0; j < len(resPath[i]); j++ {
			fmt.Print(resPath[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
