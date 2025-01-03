package main

func removeNthFromEnd(head *LinkNode, n int) *LinkNode {
	if head == nil {
		return nil
	}
	dummy := &LinkNode{}
	dummy.Next = head
	last := dummy
	pre := dummy
	//del := head
	//为了方便删除节点，走n+1步，这样pre节点的next就是要删除的节点
	for i := 0; i < n; i++ {
		last = last.Next
	}
	for last.Next != nil {
		last = last.Next
		pre = pre.Next
		//del = del.Next
	}
	pre.Next = pre.Next.Next
	return head
}
