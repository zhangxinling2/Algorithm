package main

func isPalindrome(head *ListNode) bool {
	len := 0
	node := head
	for node != nil {
		len++
		node = node.Next
	}
	if len == 1 {
		return true
	}
	node = head
	halfLen := len / 2
	newHead := &ListNode{}
	for i := 0; i < halfLen-1; i++ {
		node = node.Next
	}
	tmp := node
	node = node.Next
	tmp.Next = nil
	for node != nil {
		tmp = node
		node = node.Next
		tmp.Next = newHead.Next
		newHead.Next = tmp
	}
	node = newHead.Next
	firstNode := head
	for firstNode != nil {
		if node.Val != firstNode.Val {
			return false
		}
		node = node.Next
		firstNode = firstNode.Next
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	mid := findMid(head)
	head2 := reverseLink(mid)
	for head2 != nil {
		if head.Val != head2.Val { // 不是回文链表
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true

}

func reverseLink(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt.Next
	}
	return cur
}

func findMid(head *ListNode) *ListNode {
	pre := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		pre = pre.Next
		fast = fast.Next.Next
	}
	return pre
}
