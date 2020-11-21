package Week_03

import (
	"container/list"
)

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func permute(nums []int) [][]int {
	if len(nums) <= 0 {
		return [][]int{}
	}

	size := len(nums)
	res := make([][]int, 0, 10)
	dq := list.New()
	used := make([]uint8, size)

	permuteDfs(dq, &res, used, nums, size)
	return res
}

func permuteDfs(dq *list.List, res *[][]int, used []uint8, nums []int, size int) {
	if dq.Len() == size {
		out := make([]int, 0, size)
		element := dq.Front()
		for element != nil {
			out = append(out, element.Value.(int))
			element = element.Next()
		}
		*res = append(*res, out)

		return
	}

	for i := 0; i < size; i++ {
		if used[i] == 0 {
			dq.PushBack(nums[i])
			used[i] = 1
			permuteDfs(dq, res, used, nums, size)
			dq.Remove(dq.Back())
			used[i] = 0
		}
	}
}
