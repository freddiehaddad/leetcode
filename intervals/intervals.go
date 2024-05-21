package intervals

import "slices"

// 56. Merge Intervals
//
// Given an array of intervals where intervals[i] = [starti, endi], merge all
// overlapping intervals, and return an array of the non-overlapping intervals
// that cover all the intervals in the input.
//
// Example 1:
//
// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
//
// Example 2:
//
// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.
//
// Constraints:
//
//  1. 1 <= intervals.length <= 104
//  2. intervals[i].length == 2
//  3. 0 <= starti <= endi <= 104
func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] == b[0] {
			if a[1] == b[1] {
				return 0
			}

			if a[1] < b[1] {
				return -1
			}

			return 1
		}

		if a[0] < b[0] {
			return -1
		}

		return 1
	})

	merged := [][]int{intervals[0]}
	j := 0
	for i := 1; i < len(intervals); i++ {
		if merged[j][1] < intervals[i][0] {
			merged = append(merged, intervals[i])
			j++
		} else {
			merged[j][1] = max(merged[j][1], intervals[i][1])
		}
	}

	return merged
}
