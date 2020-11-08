package Week_01

func MoveZeros(nums []int) {
	zero := 0

	for i := 0; i < len(nums); i++ {
		// i 和 zero 同时前进，若遇到为0的情况zero就会暂停移动
		// 在下个位置遇到不为0，i ≠ zero，把不为0的元素放到zero
		// 位置
		if nums[i] != 0 {
			if i != zero {
				nums[i], nums[zero] = nums[zero], nums[i]
			}

			// 注意，zero 只要不为0就需要前进
			zero += 1
		}
	}
}
