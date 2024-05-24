package graph

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		courses       int
		prerequisites [][]int
		expected      bool
	}{
		{
			2,
			[][]int{{1, 0}},
			true,
		},
		{
			2,
			[][]int{{1, 0}, {0, 1}},
			false,
		},
		{
			5,
			[][]int{
				{0, 1},
				{0, 4},
				{0, 2},
				{1, 3},
				{2, 3},
				{4, 1},
				{4, 2},
			},
			true,
		},
	}

	for i, test := range tests {
		result := canFinish(test.courses, test.prerequisites)
		if test.expected != result {
			t.Errorf("[%d] result wrong. expected=%t got=%t",
				i, test.expected, result,
			)
		}
	}
}
