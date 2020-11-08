package Week_01

func RotateArray(nums []int, k int) {
	size := len(nums)
	if size == 0 {
		return
	}

	counts := 0

	// 每个元素都移动一次，所以counts限定边界
	// start 所指为每次开始
	for start := 0; counts < size; start++ {
		cur := start
		cache := nums[cur]

		for {
			// 移动 k 步之后的目标位置
			cur = (cur + k) % size

			// 目标位置元素cache缓存，并将原cache元素
			// 存放到目标位置
			cache, nums[cur] = nums[cur], cache
			counts += 1

			// 跳到原位置时，结束本轮调换
			// 不然往复就在那几个元素上进
			// 行
			if start == cur {
				break
			}
		}
	}
}
