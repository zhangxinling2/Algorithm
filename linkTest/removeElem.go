package main

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func removeElem(l *LinkNode, val int) *LinkNode {
	dumy := &LinkNode{}
	dumy.Next = l
	cur := dumy
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dumy.Next
}
