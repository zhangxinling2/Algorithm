package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		return
	}
	trace.Start(f)
	defer trace.Stop()
	pre := &LinkNode{
		Val: 1,
	}
	start := pre
	for i := 0; i < 10; i++ {
		h := &LinkNode{Val: i}
		pre.Next = h
		pre = pre.Next
	}
	//删除元素
	//fmt.Println(removeElem(start, 1))
	//删除倒数第N个元素
	removeNthFromEnd(start, 4)
	for start != nil {
		fmt.Println(start.Val)
		start = start.Next
	}

	//构建链表
	//m := Construct()
	//m.addAtHead(4)
	//m.addAtHead(3)
	//m.addAtHead(2)
	//m.addAtHead(1)
	////m.deleteAtIndex(0)
	//swapPairs2(m.Next)
	//for m.Next != nil {
	//	fmt.Println(m.Next.Val)
	//	m = m.Next
	//}
	//fmt.Println(m.get(2))
}
