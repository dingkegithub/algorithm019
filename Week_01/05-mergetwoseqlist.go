package Week_01

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoSeqList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	head := &ListNode{}
	tmp := head

	for !(l1 == nil && l2 == nil) {
		if l2 == nil || (l1 != nil && l1.Val < l2.Val) {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}

	return head.Next
}
