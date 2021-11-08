package tools

import "fmt"

// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func Ints2NewListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	res := new(ListNode)
	cur := res
	for _, val := range arr {
		cur.Next = &ListNode{Val: val, Next: nil}
		cur = cur.Next
	}

	return res.Next
}

func NewListNode2Ints(node *ListNode) []int {
	if node == nil {
		return nil
	}

	// 防止出现循环链表
	limit := 100
	count := 0

	arr := []int{}

	for node != nil {
		count++
		if count > limit {
			panic(fmt.Sprintf("为防止出现循环链表，限制了 ListNode 深度为 %d，请检查是否是循环链表或修改 limit 的限制。", limit))
		}
		arr = append(arr, node.Val)
		node = node.Next
	}

	return arr
}
