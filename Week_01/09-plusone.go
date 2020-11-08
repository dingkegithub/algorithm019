package Week_01

func PlusOne(nums []int) []int {
	for i := len(nums) - 1; i > 0; i-- {

		// 9 表示+1会进位，进位都是1，所以当前9
		// 的位置设为0即可
		if nums[i] == 9 {
			nums[i] = 0
		} else {
			// 不是9，则+1 不会进位，所以+1 直接返回
			nums[i] += 1
			return nums
		}
	}

	// 走到这里，说明数组[9, 9,9]类似，每一位都产生了
	// 进位，因此添加1位，高位0设置为1即可
	out := make([]int, len(nums)+1)
	out[0] = 1
	return out
}
