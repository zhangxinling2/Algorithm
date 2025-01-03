package main

func detectCycle(head *LinkNode) *LinkNode {
	stap1, stap2 := head, head
	for stap2 != nil && stap2.Next != nil {
		stap2 = stap2.Next.Next
		stap1 = stap1.Next
		if stap2 == stap1 {
			for stap1 != head {
				head = head.Next
				stap1 = stap1.Next
			}
			return head
		}
	}

	return nil
}
