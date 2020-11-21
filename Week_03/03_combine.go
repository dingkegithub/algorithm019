package Week_03

import "container/list"

/*

给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

func combine(n, k int) [][]int {
	if n < k || k <= 0 {
		return [][]int{}
	}

	dq := list.New()
	res := make([][]int, 10)
	combineDfs(dq, &res, 1, n, k)
	return res
}

func combineDfs(dq *list.List, res *[][]int, start, n, k int) {
	if dq.Len() == k {
		out := make([]int, 0, k)
		element := dq.Front()
		for element != nil {
			out = append(out, element.Value.(int))
			element = element.Next()
		}
		*res = append(*res, out)

		return
	}

	for i := start; i <= n+dq.Len()-k+1; i++ {
		dq.PushBack(i)
		combineDfs(dq, res, i+1, n, k)
		dq.Remove(dq.Back())
	}
}
