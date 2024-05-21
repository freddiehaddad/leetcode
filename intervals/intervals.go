package intervals

import "slices"

// 435. Non-overlapping Intervals
//
// Given an array of intervals intervals where intervals[i] = [starti, endi],
// return the minimum number of intervals you need to remove to make the rest
// of the intervals non-overlapping.
//
// Example 1:
//
// Input: intervals = [[1,2],[2,3],[3,4],[1,3]] Output: 1 Explanation: [1,3]
// can be removed and the rest of the intervals are non-overlapping.
//
// Example 2:
//
// Input: intervals = [[1,2],[1,2],[1,2]] Output: 2 Explanation: You need to
// remove two [1,2] to make the rest of the intervals non-overlapping.
//
// Example 3:
//
// Input: intervals = [[1,2],[2,3]] Output: 0 Explanation: You don't need to
// remove any of the intervals since they're already non-overlapping.
//
// Constraints:
//
//  1. 1 <= intervals.length <= 105
//  2. intervals[i].length == 2
//  3. -5 * 104 <= starti < endi <= 5 * 104
func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})
	var erase int
	var previous int
	for i := 1; i < len(intervals); i++ {
		if intervals[previous][1] <= intervals[i][0] {
			previous = i
		} else {
			erase++
			if intervals[i][1] <= intervals[previous][1] {
				previous = i
			}
		}
	}
	return erase
}

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
