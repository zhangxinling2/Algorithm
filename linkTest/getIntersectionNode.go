package main

//	func getIntersectionNode(headA *LinkNode, headB *LinkNode) *LinkNode {
//		m, n := 0, 0
//		curA, curB := headA, headB
//		for curA != nil {
//			m++
//			curA = curA.Next
//		}
//		for curB != nil {
//			n++
//			curB = curB.Next
//		}
//		curA, curB = headA, headB
//		for m > n {
//			curA = curA.Next
//			m--
//		}
//		for m < n {
//			curB = curB.Next
//			n--
//		}
//		for m > 0 {
//			if curA == curB {
//				return curA
//			}
//			curA = curA.Next
//			curB = curB.Next
//			m--
//		}
//		return nil
//	}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	m, n := 0, 0
	tmpA, tmpB := headA, headB
	for tmpA != nil {
		m++
		tmpA = tmpA.Next
	}
	for tmpB != nil {
		n++
		tmpB = tmpB.Next
	}
	tmpA, tmpB = headA, headB
	for m > n {
		tmpA = tmpA.Next
		m--
	}
	for m < n {
		tmpB = tmpB.Next
		n--
	}
	for m > 0 {
		if tmpA == tmpB {
			return tmpA
		} else {
			tmpA = tmpA.Next
			tmpB = tmpB.Next
			m--
		}
	}
	return nil
}
