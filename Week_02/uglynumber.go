package Week_02

import "fmt"

/*
    a b c
1 : 0 0 0 => 2 3 5 : 2         [1, 2]
2 : 1 0 0 => 4 3 5 : 3         [1, 2, 3]
3 : 1 1 0 => 4 6 5 : 4         [1, 2, 3, 4]
4 : 2 1 0 => 6 6 5 : 5         [1, 2, 3, 4, 5]
5 : 2 1 1 => 6 6 10 : 6        [1, 2, 3, 4, 5, 6]
6 : 3 2 1 => 8 9 10 : 8        [1, 2, 3, 4, 5, 6, 8]
7 : 4 2 1 => 10 9 10 : 9       [1, 2, 3, 4, 5, 6, 8, 9]
8 : 4 3 1 => 10 12 10 : 19     [1, 2, 3, 4, 5, 6, 8, 9, 10]
9 : 5 3 2 => 12 12 15 : 12     [1, 2, 3, 4, 5, 6, 8, 9, 10, 12]
*/
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	record := make([]int, n)
	record[0] = 1

	a, b, c := 0, 0, 0

	for i := 1; i < n; i++ {
		av, bv, cv := record[a]*2, record[b]*3, record[c]*5
		record[i] = min(av, bv, cv)
		fmt.Println(i, ":", a, b, c, "=>", av, bv, cv, ":", record[i])

		if record[i] == av {
			a += 1
		}
		if record[i] == bv {
			b += 1
		}
		if record[i] == cv {
			c += 1
		}
	}
	for _, v := range record {
		fmt.Print(v, " ")
	}
	fmt.Println()

	return record[len(record)-1]
}

func min(a, b, c int) int {
	if a <= b {
		if a < c {
			return a
		}
		return c
	}

	if b < c {
		return b
	}
	return c
}
