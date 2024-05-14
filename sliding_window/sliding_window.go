package slidingwindow

// 424. Longest Repeating Character Replacement
//
// You are given a string s and an integer k. You can choose any character of
// the string and change it to any other uppercase English character. You can
// perform this operation at most k times.
//
// Return the length of the longest substring containing the same letter you
// can get after performing the above operations.
//
// Constraints:
//
//  1. 1 <= s.length <= 105
//  2. s consists of only uppercase English letters.
//  3. 0 <= k <= s.length
func characterReplacement(s string, k int) int {
	histogram := [26]int{}
	var left, longest, maxFrequency int

	for right := range s {
		ch := s[right] - 'A'
		histogram[ch] += 1
		maxFrequency = max(maxFrequency, histogram[ch])
		length := right - left + 1
		if length-maxFrequency > k {
			ch = s[left] - 'A'
			histogram[ch] -= 1
			left += 1
		} else {
			longest = max(longest, length)
		}
	}

	return longest
}
