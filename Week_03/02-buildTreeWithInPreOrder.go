package Week_03

/*
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 0 || len(inorder) < 0 {
		return nil
	}
	idx := 0
	size := len(inorder)

	inorderIndex := make(map[int]int, size)
	for i, v := range inorder {
		inorderIndex[v] = i
	}

	return buildWithInPreOrder(preorder, inorder, 0, size-1, &idx, inorderIndex)
}

func buildWithInPreOrder(preorder []int, inorder []int, start, end int, idx *int, inorderIndex map[int]int) *TreeNode {
	if start < end {
		return nil
	}

	node := &TreeNode{
		Val: preorder[*idx],
	}
	*idx += 1

	if start == end {
		return node
	}

	nodePos := inorderIndex[node.Val]

	node.Left = buildWithInPreOrder(preorder, inorder, start, nodePos-1, idx, inorderIndex)
	node.Left.Right = buildWithInPreOrder(preorder, inorder, nodePos+1, end, idx, inorderIndex)

	return node
}
