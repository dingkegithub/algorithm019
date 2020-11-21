package Week_03

import (
	"container/list"
	"sort"
)

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。



示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]


提示：

1 <= nums.length <= 8
-10 <= nums[i] <= 10

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func permute2(nums []int) [][]int {
	if len(nums) <= 0 {
		return [][]int{}
	}

	size := len(nums)
	res := make([][]int, 0, 10)
	dq := list.New()
	used := make([]uint8, size)

	sort.Ints(nums)
	permuteDfs(dq, &res, used, nums, size)
	return res
}

func permuteDfs2(dq *list.List, res *[][]int, used []uint8, nums []int, size int) {
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
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 {
				continue
			}

			dq.PushBack(nums[i])
			used[i] = 1
			permuteDfs2(dq, res, used, nums, size)
			dq.Remove(dq.Back())
			used[i] = 0
		}
	}
}
