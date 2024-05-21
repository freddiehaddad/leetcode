package intervals

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
	}

	for i, test := range tests {
		actual := merge(test.input)
		if slices.CompareFunc(
			test.expected, actual, compareIntervals,
		) != 0 {
			t.Errorf("[%d] result wrong. expected=%v got=%v",
				i, test.expected, actual,
			)
		}
	}
}

func compareIntervals(s1, s2 []int) int {
	return slices.Compare(s1, s2)
}
