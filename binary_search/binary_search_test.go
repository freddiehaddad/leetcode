package binarysearch

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		input_piles []int
		input_hours int
		expected    int
	}{
		{
			[]int{3, 6, 7, 11},
			8,
			4,
		},
		{
			[]int{30, 11, 23, 4, 20},
			5,
			30,
		},
		{
			[]int{30, 11, 23, 4, 20},
			6,
			23,
		},
		{
			[]int{312884470},
			968709470,
			1,
		},
		{
			[]int{312884470},
			312884469,
			2,
		},
		{
			[]int{
				831235932,
				437082868,
				576572631,
				529869153,
				55330371,
				511060323,
				581115062,
				931692072,
				600856556,
				519045509,
				504164418,
				431105822,
				580257183,
			},
			964065706,
			8,
		},
	}

	for i, test := range tests {
		actual := minEatingSpeed(test.input_piles, test.input_hours)
		if test.expected != actual {
			t.Errorf("[%d] result wrong. expected=%d got=%d",
				i, test.expected, actual,
			)
		}
	}
}
