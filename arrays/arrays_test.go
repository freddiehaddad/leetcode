package arrays

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			[]int{1},
			1,
			[]int{},
			0,
			[]int{1},
		},
		{
			[]int{0, 0},
			1,
			[]int{1},
			1,
			[]int{0, 1},
		},
	}

	for i, test := range tests {
		merge(test.nums1, test.m, test.nums2, test.n)
		if slices.Compare(test.nums1, test.expected) != 0 {
			t.Errorf("[%d] result wrong. expected=%v got=%v",
				i, test.expected, test.nums1)
		}
	}
}
