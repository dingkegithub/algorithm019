package Week_02

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
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

}
