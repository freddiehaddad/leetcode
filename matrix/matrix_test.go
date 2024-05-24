package matrix

import (
	"slices"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected []int
	}{
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			[]int{
				1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7,
				11, 10,
			},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}

	for i, test := range tests {
		result := spiralOrder(test.input)
		if slices.Compare(test.expected, result) != 0 {
			t.Errorf("[%d] result wrong. expected=%v got=%v",
				i, test.expected, result,
			)
		}
	}
}
