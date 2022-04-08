package tree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		l := make([]int, size)
		for i := 0; i < size; i++ {
			l[i] = queue[0].Val
			for _, child := range queue[0].Children {
				queue = append(queue, child)
			}
			queue = queue[1:]
		}
		res = append(res, l)
	}
	return res
}
