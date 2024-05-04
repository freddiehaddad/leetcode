package arrays

import (
	"slices"
	"testing"
)

func TestMajoriyElement(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{1},
			1,
		},
		{
			[]int{1, 2, 1},
			1,
		},
		{
			[]int{2, 1, 2},
			2,
		},
		{
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},
	}

	for i, test := range tests {
		m := majorityElement(test.input)
		if test.expected != m {
			t.Errorf("[%d] result wrong. expected=%d got=%d",
				i, test.expected, m)
		}
	}
}

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

func TestRotateArray(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{
			[]int{1},
			0,
			[]int{1},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
		{
			[]int{1},
			2,
			[]int{1},
		},
		{
			[]int{1, 2},
			0,
			[]int{1, 2},
		},
		{
			[]int{1, 2},
			1,
			[]int{2, 1},
		},
		{
			[]int{1, 2},
			2,
			[]int{1, 2},
		},
		{
			[]int{1, 2},
			3,
			[]int{2, 1},
		},
		{
			[]int{1, 2},
			4,
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3},
			0,
			[]int{1, 2, 3},
		},
		{
			[]int{1, 2, 3},
			1,
			[]int{3, 1, 2},
		},
		{
			[]int{1, 2, 3},
			2,
			[]int{2, 3, 1},
		},
		{
			[]int{1, 2, 3},
			3,
			[]int{1, 2, 3},
		},
		{
			[]int{1, 2, 3},
			4,
			[]int{3, 1, 2},
		},
		{
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
	}

	for i, test := range tests {
		rotate(test.input, test.k)
		if slices.Compare(
			test.input, test.expected) != 0 {
			t.Errorf("[%d] result wrong. expected=%v got=%v",
				i, test.expected, test.input)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			[]int{0},
			1,
			[]int{0},
		},
		{
			[]int{1, 1, 2},
			2,
			[]int{1, 2},
		},
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
			[]int{0, 1, 2, 3, 4},
		},
	}

	for i, test := range tests {
		k := removeDuplicates(test.nums)
		if test.k != k {
			t.Errorf("[%d] length wrong. expected=%d got=%d",
				i, test.k, k)
		}
		if slices.Compare(
			test.nums[:test.k], test.expected) != 0 {
			t.Errorf("[%d] result wrong. expected=%v got=%v",
				i, test.expected, test.nums[:test.k])
		}
	}
}

func TestRemoveDuplicatesII(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			[]int{0},
			1,
			[]int{0},
		},
		{
			[]int{0, 0},
			2,
			[]int{0, 0},
		},
		{
			[]int{0, 1},
			2,
			[]int{0, 1},
		},
		{
			[]int{1, 1, 2},
			3,
			[]int{1, 1, 2},
		},
		{
			[]int{1, 2, 2},
			3,
			[]int{1, 2, 2},
		},
		{
			[]int{1, 1, 1, 2},
			3,
			[]int{1, 1, 2},
		},
		{
			[]int{1, 2, 2, 2},
			3,
			[]int{1, 2, 2},
		},
		{
			[]int{1, 1, 1, 2, 2, 3},
			5,
			[]int{1, 1, 2, 2, 3},
		},
		{
			[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			7,
			[]int{0, 0, 1, 1, 2, 3, 3},
		},
	}

	for i, test := range tests {
		k := removeDuplicatesII(test.nums)
		if test.k != k {
			t.Errorf("[%d] length wrong. expected=%d got=%d",
				i, test.k, k)
		}
		if slices.Compare(
			test.nums[:test.k], test.expected) != 0 {
			t.Errorf("[%d] result wrong. expected=%v got=%v",
				i, test.expected, test.nums[:test.k])
		}
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		expected []int
	}{
		{
			[]int{1},
			1,
			[]int{},
		},
		{
			[]int{3, 2, 2, 3},
			3,
			[]int{2, 2},
		},
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			[]int{0, 1, 3, 0, 4},
		},
	}

	for i, test := range tests {
		k := removeElement(test.nums, test.val)
		if len(test.expected) != k {
			t.Errorf("[%d] length wrong. expected=%d got=%d",
				i, len(test.expected), k)
		}
		if slices.Compare(
			test.nums[:len(test.expected)], test.expected) != 0 {
			t.Errorf("[%d] result wrong. expected=%v got=%v",
				i, test.expected, test.nums[:len(test.expected)])
		}
	}
}
