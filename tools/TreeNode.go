package tools

// 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 输入数组数据中用 NULL 代替 nil
var NULL = -1 << 63

func Ints2NewTreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: ints[0]}

	// 创建一个 length 1 capacity n * 2 的 slice
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		// 从队列中取出并移除
		node := queue[0]
		queue = queue[1:]

		// 此时 queue 是空队列
		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			// 将左节点放入队列
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {

			node.Right = &TreeNode{Val: ints[i]}
			// 将右节点放入队列
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
