package randomizedset

import "testing"

func TestRandomizedSet(t *testing.T) {
	rs := Constructor()
	if !rs.Insert(1) {
		t.Errorf("Insert(1): expected=%t got=%t", true, false)
	}

	if rs.Remove(2) {
		t.Errorf("Remove(2): expected=%t got=%t", false, true)
	}

	if !rs.Insert(2) {
		t.Errorf("Insert(2): expected=%t got=%t", true, false)
	}

	val := rs.GetRandom()
	if val != 1 && val != 2 {
		t.Errorf("GetRandom(): expected={%d, %d} got=%d", 1, 2, val)
	}

	if !rs.Remove(1) {
		t.Errorf("Remove(1): expected=%t got=%t", true, false)
	}

	if rs.Insert(2) {
		t.Errorf("Insert(2): expected=%t got=%t", false, true)
	}

	val = rs.GetRandom()
	if val != 2 {
		t.Errorf("GetRandom(): expected=%d got=%d", 2, val)
	}
}
