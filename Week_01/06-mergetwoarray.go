package Week_01

func MergeTwoSeqArray(nums1 []int, m int, nums2 []int, n int) {
	// 指向两个数组的最后一个元素
	nums1P, nums2P := m-1, n-1

	// 从nums最后开始存放元素
	for k := m + n - 1; k >= 0; k-- {
		// 归并排序
		if nums2P < 0 || (nums1P >= 0 && nums1[nums1P] > nums2[nums2P]) {
			nums1[k] = nums1[nums1P]
			nums1P -= 1
		} else {
			nums1[k] = nums2[nums2P]
			nums2P -= 1
		}
	}
}
