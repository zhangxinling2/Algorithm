package main

func swapPairs(head *MyLink) *MyLink {
	h := &MyLink{
		val:  0,
		next: head,
	}
	pre := h
	cur := head
	for cur != nil && cur.next != nil {
		next := cur.next
		pre.next = next
		cur.next = next.next
		next.next = cur
		pre = cur
		cur = cur.next
	}
	return h.next
}

// swapPairs2 递归法
func swapPairs2(head *MyLink) *MyLink {
	if head == nil || head.next == nil {
		return head
	}
	next := head.next
	//head的next就是交换后的后续节点,需要交换的是next.Next，当前的next需要和head互换
	head.next = swapPairs(next.next)
	next.next = head
	return next
}
