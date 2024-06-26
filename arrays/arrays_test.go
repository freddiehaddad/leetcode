package arrays

import (
	"slices"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		gas      []int
		cost     []int
		expected int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			3,
		},
		{
			[]int{2, 3, 4},
			[]int{3, 4, 3},
			-1,
		},
	}

	for i, test := range tests {
		start := canCompleteCircuit(test.gas, test.cost)
		if test.expected != start {
			t.Errorf("[%d] result wrong. expected=%d got=%d",
				i, test.expected, start)
		}
	}
}

func TestCandy(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{1, 2, 2},
			4,
		},
		{
			[]int{1, 0, 2},
			5,
		},
		{
			[]int{1, 3, 2, 2, 1},
			7,
		},
	}

	for i, test := range tests {
		count := candy(test.input)
		if test.expected != count {
			t.Errorf("[%d] result wrong. expected=%d got=%d",
				i, test.expected, count)
		}
	}
}

func TestCanJump(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			[]int{1},
			true,
		},
		{
			[]int{0},
			true,
		},
		{
			[]int{1, 0},
			true,
		},
		{
			[]int{0, 1},
			false,
		},
		{
			[]int{2, 3, 1, 1, 4},
			true,
		},
		{
			[]int{3, 2, 1, 0, 4},
			false,
		},
	}

	for i, test := range tests {
		m := canJump(test.input)
		if test.expected != m {
			t.Errorf("[%d] result wrong. expected=%t got=%t",
				i, test.expected, m)
		}
	}
}

func TestCanJumpII(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{1},
			0,
		},
		{
			[]int{0},
			0,
		},
		{
			[]int{1, 0},
			1,
		},
		{
			[]int{2, 3, 1, 1, 4},
			2,
		},
		{
			[]int{2, 3, 0, 1, 4},
			2,
		},
	}

	for i, test := range tests {
		jumps := canJumpII(test.input)
		if test.expected != jumps {
			t.Errorf("[%d] result wrong. expected=%d got=%d",
				i, test.expected, jumps)
		}
	}
}

func TestHIndex(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{1, 3, 1},
			1,
		},
		{
			[]int{3, 0, 6, 1, 5},
			3,
		},
	}

	for i, test := range tests {
		result := hIndex(test.input)
		if test.expected != result {
			t.Errorf("[%d] result wrong. expected=%d got=%d",
				i, test.expected, result)
		}
	}
}

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

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{10},
			0,
		},
		{
			[]int{1, 3},
			2,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
	}

	for i, test := range tests {
		m := maxProfit(test.input)
		if test.expected != m {
			t.Errorf("[%d] result wrong. expected=%d got=%d",
				i, test.expected, m)
		}
	}
}

func TestMaxProfitII(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{10},
			0,
		},
		{
			[]int{1, 3},
			2,
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
		{
			[]int{7, 1, 5, 3, 6, 4},
			7,
		},
	}

	for i, test := range tests {
		m := maxProfitII(test.input)
		if test.expected != m {
			t.Errorf("[%d] result wrong. expected=%d got=%d",
				i, test.expected, m)
		}
	}
}

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
	}

	for i, test := range tests {
		result := maxSubArray(test.input)
		if test.expected != result {
			t.Errorf("[%d] result wrong. expected=%d got=%d",
				i, test.expected, result)
		}
	}
}

func TestMaxSubarraySumCircular(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{1, -2, 3, -2},
			3,
		},
		{
			[]int{5, -3, 5},
			10,
		},
		{
			[]int{-3, -2, -3},
			-2,
		},
	}

	for i, test := range tests {
		result := maxSubarraySumCircular(test.input)
		if test.expected != result {
			t.Errorf("[%d] result wrong. expected=%d got=%d",
				i, test.expected, result)
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

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
		{
			[]int{-1, 1, 0, -3, 3},
			[]int{0, 0, 9, 0, 0},
		},
	}

	for i, test := range tests {
		products := productExceptSelf(test.input)
		if slices.Compare(
			test.expected, products) != 0 {
			t.Errorf("[%d] result wrong. expected=%v got=%v",
				i, test.expected, products)
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

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"the sky is blue",
			"blue is sky the",
		},
		{
			"  hello world  ",
			"world hello",
		},
		{
			"a good   example",
			"example good a",
		},
	}

	for i, test := range tests {
		result := reverseWords(test.input)
		if test.expected != result {
			t.Errorf("[%d] result wrong. expected=%q got=%q",
				i, test.expected, result,
			)
		}
	}
}
