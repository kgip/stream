package impl

import "stream/utils/sort"

//定义全局变量
var Sorter sort.Sorter = &QuickSort{}

//快速排序
type QuickSort struct {
}

func (quickSort *QuickSort) DoSort(list []interface{}, comparator sort.Comparator) error {
	return quickSort.partition(list, 0, len(list)-1, comparator)
}

func (quickSort *QuickSort) partition(list []interface{}, left, right int, comparator sort.Comparator) error {
	if right-left <= 0 {
		return nil
	}
	start := left
	end := right
	isLeft := true //pivot是否在左侧
	pivot := left
	pivotValue := list[pivot]
	for left < right {
		if isLeft {
			if flag, err := comparator(pivotValue, list[right]); err != nil {
				return err
			} else if flag {
				list[pivot] = list[right]
				list[right] = pivotValue
				pivot = right
				left++
				isLeft = false
			} else {
				right--
			}
		} else {
			if flag, err := comparator(list[left], pivotValue); err != nil {
				return err
			} else if flag {
				list[pivot] = list[left]
				list[left] = pivotValue
				pivot = left
				right--
				isLeft = true
			} else {
				left++
			}
		}
	}
	if err := quickSort.partition(list, start, pivot-1, comparator); err != nil {
		return err
	}
	if err := quickSort.partition(list, pivot+1, end, comparator); err != nil {
		return err
	}
	return nil
}
