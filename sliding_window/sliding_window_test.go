package slidingwindow

import "testing"

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		input_s  string
		input_k  int
		expected int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		{"ABCBAAAA", 1, 5},
	}

	for i, test := range tests {
		result := characterReplacement(test.input_s, test.input_k)
		if test.expected != result {
			t.Errorf("[%d] result wrong. expected=%d got=%d",
				i, test.expected, result,
			)
		}
	}
}
