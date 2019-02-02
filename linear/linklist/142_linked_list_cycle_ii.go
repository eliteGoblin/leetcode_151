package linklist

func detectCycleNaive(head *ListNode) *ListNode {
	mp := make(map[*ListNode]bool)
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := mp[cur]; ok {
			return cur
		} else {
			mp[cur] = true
		}
	}
	return nil
}

func detectCycle(head *ListNode) *ListNode {
	loopLength := findLoopLength(head)
	if 0 == loopLength {
		return nil
	}
	var first, last *ListNode
	for first = head; loopLength > 0; first = first.Next {
		loopLength--
	}
	for last = head; first != last; {
		first = first.Next
		last = last.Next
	}
	return first
}

func findLoopLength(head *ListNode) int {
	if nil == head {
		return 0
	}
	slow, fast := head, head.Next
	var startPos *ListNode
	for slow != nil && fast != nil {
		if slow == fast {
			startPos = slow
			break
		}
		slow = slow.Next
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			return 0
		}
	}
	if startPos == nil {
		return 0
	}
	res := 1
	for cur := startPos.Next; cur != startPos; {
		cur = cur.Next
		res++
	}
	return res
}
