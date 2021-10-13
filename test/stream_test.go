package test

import "testing"

func TestSliceLen(t *testing.T) {
	list := make([]int,10)
	t.Log(len(list))
	for i, e := range list {
		t.Log(i,e)
	}
}
