package test

import (
	"stream/utils/sort/impl"
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := []interface{}{2,5,1,3,0,3,7,3,2,20,24}
	err := impl.Sorter.DoSort(list, func(o1, o2 interface{}) (bool, error) {
		n := o1.(int) - o2.(int)
		return n > 0, nil
	})

	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}