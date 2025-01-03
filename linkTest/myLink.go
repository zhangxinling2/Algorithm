package main

type ListNode struct {
	Val  int
	Next *ListNode
}
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
