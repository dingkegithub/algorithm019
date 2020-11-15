package Week_02

import "container/heap"

//
// [ [值: 频次],  [值: 频次],  [值: 频次],  [值: 频次] ]
//
type NumFreqSlice [][2]int

// Len is the number of elements in the collection.
func (n *NumFreqSlice) Len() int {
	return len(*n)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (n *NumFreqSlice) Less(i int, j int) bool {
	return (*n)[i][1] < (*n)[j][1]
}

// Swap swaps the elements with indexes i and j.
func (n *NumFreqSlice) Swap(i int, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}

func (n *NumFreqSlice) Push(x interface{}) {
	*n = append(*n, x.([2]int))
}

func (n *NumFreqSlice) Pop() interface{} {
	size := len(*n)

	popItem := (*n)[size-1]
	*n = (*n)[:size-1]
	return popItem
}

func TopKFrequent(nums []int, k int) []int {
	if len(nums) <=0 || k == 0 {
		return []int{}
	}

	m := make(map[int]int)

	//
	// 统计每个值的频次
	//
	for _, val := range nums {
		m[val]++
	}

	nfs := &NumFreqSlice{}
	heap.Init(nfs)

	//
	// 入堆，固定堆大小为 k
	//
	for key, val := range m {
		heap.Push(nfs, [2]int{key, val})
		for nfs.Len() > k {
			heap.Pop(nfs)
		}
	}

	//
	// 取出堆中的元素
	//
	res := make([]int, k)
	for i:=0; i<k; i++ {
		res[i] = heap.Pop(nfs).([2]int)[0]
	}

	return res
}
