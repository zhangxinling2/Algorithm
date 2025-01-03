package main

func ReverseLink(head *LinkNode) *LinkNode {
	if head == nil {
		return nil
	}
	m := &LinkNode{
		Val:  0,
		Next: nil,
	}
	for head != nil {
		t := head
		head = head.Next
		t.Next = m.Next
		m.Next = t
	}
	return m.Next
}
func ReverseLink2(head *LinkNode) *LinkNode {
	if head == nil {
		return nil
	}
	var pre *LinkNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
func ReverseLink3(head *LinkNode) *LinkNode {
	return helpReverse(nil, head)
}
func helpReverse(pre *LinkNode, cur *LinkNode) *LinkNode {
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	return helpReverse(cur, next)
}
