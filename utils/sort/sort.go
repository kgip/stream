package sort

type Comparator func(o1,o2 interface{}) (bool,error)

//排序接口
type Sorter interface {
	DoSort(list []interface{},comparator Comparator) error
}