package ds

import "testing"

func TestSet(t *testing.T) {
	set := NewSet(10)
	set.Add(10)
	if set.Add(10) {
		t.Fail()
	}
	set.Add(20)
	set.Add(20)
	if set.Size() != 2 {
		t.Fail()
	}
	if !set.IsExist(10) {
		t.Fail()
	}
	if set.IsExist(30) {
		t.Fail()
	}
	set.Remove(10)
	set.Remove(20)
	if set.Size() != 0 {
		t.Fail()
	}
}




