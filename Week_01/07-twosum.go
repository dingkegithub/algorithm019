package Week_01

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, item := range nums {
		if v, ok := m[target-item]; ok {
			return []int{v, i}
		}
		m[item] = i
	}
	return nil
}
