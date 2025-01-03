package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	f, s := head, head
	var ps *ListNode
	for f != nil && f.Next != nil {
		f = f.Next.Next
		ps = s
		s = s.Next
	}
	l := sortList(head)
	r := sortList(ps)
	return Merge(l, r)
}
func Merge(l, r *ListNode) *ListNode {
	var dummy = &ListNode{}
	var curr = dummy
	for l != nil && r != nil {
		if l.Val < r.Val {
			curr.Next = l
			l = l.Next
		} else {
			curr.Next = r
			r = r.Next
		}
		curr = curr.Next
	}
	if l == nil {
		curr.Next = r
	} else {
		curr.Next = l
	}
	return dummy.Next
}
