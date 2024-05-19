package timebasedkvstore

import (
	"slices"
	"testing"
)

func TestSet(t *testing.T) {
	tests := []struct {
		key    string
		values []TimeValue
	}{
		{
			"a",
			[]TimeValue{
				{"0", 0},
				{"1", 1},
			},
		},
		{
			"b",
			[]TimeValue{
				{"2", -1},
				{"3", 0},
			},
		},
	}

	tm := Constructor()
	for _, test := range tests {
		for _, value := range test.values {
			tm.Set(test.key, value.Value, value.Timestamp)
		}
	}

	for i, test := range tests {
		values, ok := tm.store[test.key]
		if !ok {
			t.Errorf("[%d] entries wrong. expected=%v got=%v",
				i, test.values, values)
			continue
		}
		if slices.CompareFunc(
			test.values,
			values,
			func(s1, s2 TimeValue) int {
				if s1 == s2 {
					return 0
				}

				if s1.Value < s2.Value {
					return -1
				}

				if s1.Timestamp < s2.Timestamp {
					return -1
				}

				return 1
			},
		) != 0 {
			t.Errorf(
				"[%d] values wrong. expected=%v got=%v",
				i, test.values, values,
			)
		}
	}
}

func TestGet(t *testing.T) {
	testData := []struct {
		key    string
		values []TimeValue
	}{
		{
			"a",
			[]TimeValue{
				{"0", 0},
				{"1", 10},
			},
		},
		{
			"b",
			[]TimeValue{
				{"2", -10},
				{"3", 0},
			},
		},
	}

	tm := Constructor()
	for _, data := range testData {
		for _, value := range data.values {
			tm.Set(data.key, value.Value, value.Timestamp)
		}
	}

	tests := []struct {
		key       string
		timestamp int
		expected  string
	}{
		{"a", -1, ""},
		{"a", 0, "0"},
		{"a", 1, "0"},
		{"a", 9, "0"},
		{"a", 10, "1"},
		{"a", 11, "1"},
		{"a", 12, "1"},
		{"b", -11, ""},
		{"b", -10, "2"},
		{"b", -9, "2"},
		{"b", -1, "2"},
		{"b", 0, "3"},
		{"b", 1, "3"},
		{"b", 2, "3"},
	}

	for i, test := range tests {
		actual := tm.Get(test.key, test.timestamp)
		if test.expected != actual {
			t.Errorf("[%d] result wrong. expected=%s got=%s",
				i, test.expected, actual)
		}
	}
}
