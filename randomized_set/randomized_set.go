package randomizedset

import "math/rand"

// Implement the RandomizedSet class:
//
//   - RandomizedSet() Initializes the RandomizedSet object.
//
//   - bool insert(int val) Inserts an item val into the set if not
//     present. Returns true if the item was not present, false
//     otherwise.
//
//   - bool remove(int val) Removes an item val from the set if
//     present. Returns true if the item was present, false otherwise.
//
//   - int getRandom() Returns a random element from the current set
//     of elements (it's guaranteed that at least one element exists
//     when this method is called). Each element must have the same
//     probability of being returned.
//
// You must implement the functions of the class such that each function works
// in average O(1) time complexity.
type RandomizedSet struct {
	data    map[int]int
	indices []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		data:    make(map[int]int),
		indices: make([]int, 0),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	// exists?
	if _, ok := rs.data[val]; ok {
		return false
	}

	// append the element to indices array and
	// record its index in the map
	i := len(rs.indices)
	rs.indices = append(rs.indices, val)
	rs.data[val] = i
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	// exists?
	i, ok := rs.data[val]
	if !ok {
		return false
	}

	delete(rs.data, val)

	// last index?
	if i == len(rs.indices)-1 {
		rs.indices = rs.indices[:i]
		return true
	}

	// middle index
	last := len(rs.indices) - 1
	v := rs.indices[last]
	rs.indices[i] = v
	rs.data[v] = i
	rs.indices = rs.indices[:last]
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	n := len(rs.indices)
	i := rand.Intn(n)
	return rs.indices[i]
}
