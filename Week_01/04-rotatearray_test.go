package Week_01

import "testing"

func TestRotateArray(t *testing.T) {

	in := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{1, 2},
		{1},
	}

	wantK3 := [][]int{
		{3, 4, 5, 1, 2},
		{4, 5, 6, 1, 2, 3},
		{2, 1},
		{1},
	}

	for j, a := range in {
		RotateArray(a, 3)
		for i, v := range wantK3[j] {
			if a[i] != v {
				t.Error("K3 want array[", i, "]=", v, "but: ", v)
				t.FailNow()
			}
		}
	}
}
