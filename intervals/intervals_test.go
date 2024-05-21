package intervals

import (
	"slices"
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			1,
		},
		{
			[][]int{{1, 2}, {1, 2}, {1, 2}},
			2,
		},
		{
			[][]int{{1, 2}, {2, 3}},
			0,
		},
		{
			[][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}},
			2,
		},
	}

	for i, test := range tests {
		result := eraseOverlapIntervals(test.input)
		if test.expected != result {
			t.Errorf("[%d] result wrong. expected=%d got=%d",
				i, test.expected, result,
			)
		}
	}
}

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
