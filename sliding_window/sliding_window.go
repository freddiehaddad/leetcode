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

// 567. Permutation in String
//
// Given two strings s1 and s2, return true if s2 contains a permutation of s1,
// or false otherwise.
//
// In other words, return true if one of s1's permutations is the substring of
// s2.
//
// Constraints:
//
//  1. 1 <= s1.length, s2.length <= 104
//  2. s1 and s2 consist of lowercase English letters.
func checkInclusion(s1 string, s2 string) bool {
	histogram := [26]int{}
	// tally indicies part of s1
	for i := range s1 {
		ch := s1[i] - 'a'
		histogram[ch]--
	}

	// mark indicies not part of s1
	for i := range histogram {
		if histogram[i] == 0 {
			histogram[i] = 1
		}
	}

	count := len(s1)
	left := 0
	for right := range s2 {
		ch := s2[right] - 'a'
		frequency := histogram[ch]
		// current letter in s2 is in s1
		switch frequency {
		case 1: // Current letter is not a part of the string.
			// Window size is 0.
			if count == len(s1) {
				left++
				continue
			}
			// Shrink the window to 0 adding back letters.
			for left < right {
				_ch := s2[left] - 'a'
				histogram[_ch]--
				left++
				count++
			}
		case 0: // Current letter exceeds needed amount.
			// Slide window untl the earlier occurrence of the
			// letter is outside.
			_ch := s2[left] - 'a'
			for _ch != ch {
				histogram[_ch]--
				count++
				left++
				_ch = s2[left] - 'a'
			}
			// Advance to next letter.
			left++

		default: // Letter is part of string and can be counted.
			histogram[ch]++
			// Start of the window?
			if count == len(s1) {
				left = right
			}
			count--
		}

		// Did we find all the letters in s1?
		if count == 0 {
			return true
		}
	}

	return false
}
