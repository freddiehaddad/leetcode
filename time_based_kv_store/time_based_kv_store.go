package timebasedkvstore

// 981. Time Based Key-Value Store
//
// Design a time-based key-value data structure that can store multiple values
// for the same key at different time stamps and retrieve the key's value at a
// certain timestamp.
//
// Implement the TimeMap class:
//
//   1. TimeMap() Initializes the object of the data structure.
//   2. void set(String key, String value, int timestamp) Stores the key key
//      with the value value at the given time timestamp.
//   3. String get(String key, int timestamp) Returns a value such that set was
//      called previously, with timestamp_prev <= timestamp. If there are
//      multiple such values, it returns the value associated with the largest
//      timestamp_prev. If there are no values, it returns "".
//
// Constraints:
//
//   1. 1 <= key.length, value.length <= 100
//   2. key and value consist of lowercase English letters and digits.
//   3. 1 <= timestamp <= 107
//   4. All the timestamps timestamp of set are strictly increasing.
//   5. At most 2 * 105 calls will be made to set and get.

type TimeValue struct {
	Value     string
	Timestamp int
}

type timeValues []TimeValue

type TimeMap struct {
	store map[string]timeValues
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string]timeValues),
	}
}

func (t *TimeMap) Set(key, value string, timestamp int) {
	tv := TimeValue{
		value,
		timestamp,
	}

	t.store[key] = append(t.store[key], tv)
}

func (t *TimeMap) Get(key string, timestamp int) string {
	var tv *TimeValue

	values := t.store[key]
	left := 0
	right := len(values)

	for left < right {
		mid := left/2 + right/2
		_tv := &values[mid]
		if _tv.Timestamp > timestamp {
			right = mid
		} else {
			tv = _tv
			left = mid + 1
		}
	}

	if tv == nil {
		return ""
	}

	return tv.Value
}
