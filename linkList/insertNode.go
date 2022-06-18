package linkList

//给定循环单调非递减列表中的一个点，写一个函数向这个列表中插入一个新元素 insertVal ，使这个列表仍然是循环升序的。
//
//给定的可以是这个列表中任意一个顶点的指针，并不一定是这个列表中最小元素的指针。
//
//如果有多个满足条件的插入位置，可以选择任意一个位置插入新的值，插入后整个列表仍然保持有序。
//
//如果列表为空（给定的节点是 null），需要创建一个循环有序列表并返回这个节点。否则。请返回原先给定的节点。
//
//
//
//示例 1：
//
//
//
//
//输入：head = [3,4,1], insertVal = 2
//输出：[3,4,1,2]
//解释：在上图中，有一个包含三个元素的循环有序列表，你获得值为 3 的节点的指针，我们需要向表中插入元素 2 。新插入的节点应该在 1 和 3 之间，插入之后，整个列表如上图所示，最后返回节点 3 。
//
//
//示例 2：
//
//输入：head = [], insertVal = 1
//输出：[1]
//解释：列表为空（给定的节点是 null），创建一个循环有序列表并返回这个节点。
//示例 3：
//
//输入：head = [1], insertVal = 0
//输出：[1,0]
//
//
//提示：
//
//0 <= Number of Nodes <= 5 * 10^4
//-10^6 <= Node.val <= 10^6
//-10^6 <= insertVal <= 10^6
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/4ueAj6
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//Definition for a Node.
type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		return &Node{Val: x}
	}
	if aNode.Next == aNode {
		node := &Node{Val: x}
		aNode.Next = node
		node.Next = aNode
	}
	cur := aNode
	//找到<=x的最大的点
	var mx *Node
	//>=x的最大的点
	var min *Node
	for cur.Next != aNode {
		if cur.Val <= x && (mx == nil || cur.Val > mx.Val) {
			mx = cur
		}
		if cur.Val >= x && (min == nil || cur.Val > min.Val) {
			min = cur
		}
		cur = cur.Next
	}
	if mx != nil {
		node := &Node{Val: x}
		next := mx.Next
		mx.Next = node
		node.Next = next
	} else {
		node := &Node{Val: x}
		next := min.Next
		min.Next = node
		node.Next = next
	}
	return aNode
}
