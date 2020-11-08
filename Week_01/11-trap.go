package Week_01

func Trap(height []int) int {
	size := len(height)
	if size <= 0 {
		return 0
	}

	lEdge := make([]int, size)
	rEdge := make([]int, size)

	// 求出每个元素代表桶的左边沿高度
	lEdge[0] = height[0]
	for i := 1; i < size; i++ {
		lEdge[i] = max(height[i], lEdge[i-1])
	}

	// 求出每个元素代表桶的右边沿高度
	rEdge[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		rEdge[i] = max(height[i], rEdge[i+1])
	}

	capacity := 0

	for i := 1; i < size-1; i++ {
		// 每一个位置储水量是 最低的边沿 - 桶底高度
		capacity += min(rEdge[i], lEdge[i]) - height[i]
	}

	return capacity
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
