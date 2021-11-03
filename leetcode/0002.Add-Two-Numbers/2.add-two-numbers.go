package leetcode

import . "github.com/otfot/LeetCode-go/tools"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, head := 0, new(ListNode)

	for cur := head; l1 != nil || l2 != nil || carry != 0; cur = cur.Next {

		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: carry % 10, Next: nil}
		carry /= 10

	}

	return head.Next

}
