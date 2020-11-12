package Week_02

import "container/list"

//
// 递归：思路同二叉树先序遍历，根---child1-child2-child3-child4
//
func NTreePreOrderL1(root *Node) []int {
	o := make([]int, 0, 10)
	nTreePreOrderL1(root, &o)
	return o
}

func nTreePreOrderL1(node *Node, o *[]int) {
	if node == nil {
		return
	}

	*o = append(*o, node.Val)

	for _, v := range node.Children {
		nTreePreOrderL1(v, o)
	}
}

//
// 使用堆栈
//
func NTreePreOrderL2(root *Node) []int {
	if root == nil {
		return []int{}
	}

	tmp := root
	o := make([]int, 0, 10)

	stack := list.New()
	stack.PushBack(tmp)

	for stack.Len() > 0 {
		item := stack.Back()
		node := stack.Remove(item).(*Node)
		o = append(o, node.Val)

		// 注意此处先输出的后进栈
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.PushBack(node.Children[i])
		}
	}

	return o
}
