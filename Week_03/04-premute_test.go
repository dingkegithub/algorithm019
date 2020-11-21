package Week_03

import (
	"fmt"
	"testing"
)

func TestPremute01(t *testing.T) {
	in := []int{1, 2, 3}

	/*
		want := [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}
	*/
	res := permute(in)
	for i := 0; i < len(res); i++ {
		fmt.Println("[")
		fmt.Print("\t")
		for _, v := range res[i] {
			fmt.Print(v, ",")
		}
		fmt.Println()
	}
}
