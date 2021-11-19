package main

func midOrUpMidNode(head *ListNode) *ListNode {
	if head.Next == nil || head.Next.Next == nil {
		return head
	}

	slow, fast := head.Next, head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}

	return slow
}
func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	slow, fast := head.Next, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}

	return slow
}

