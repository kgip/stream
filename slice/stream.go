package slicestream

import (
	"github.com/pkg/errors"
	"stream/utils/sort"
	"stream/utils/sort/impl"
)

var (
	sorter = impl.Sorter
)

type listStream struct {
	list []interface{} //被操作的切片
	err  error         //流操作中的异常
}

// 创建一个流对象
func Stream(src interface{}) *listStream {
	if src == nil {
		return &listStream{nil,errors.New("Input param can not be nil!")}
	}
	//试图将src转换成切片类型
	list,ok := src.([]interface{})
	if !ok {
		return &listStream{nil,errors.New("Input param is not slice type")}
	}
	dst := make([]interface{}, len(list))
	//对源切边进行拷贝，避免并发问题
	copy(dst, list)
	stream := &listStream{list, nil}
	return stream
}

//元素遍历
func (stream *listStream) ForEach(action func(index int, item interface{}) error) *listStream {
	//如果之前的流操作报错，后面的流操作不予执行
	if stream.err != nil || len(stream.list) <= 0{
		return stream
	}
	for i, e := range stream.list {
		err := action(i, e)
		if err != nil {
			stream.err = errors.Wrap(err, "Failed at ForEach:")
			break
		}
	}
	return stream
}

//元素过滤
func (stream *listStream) Filter(filter func(index int, item interface{}) (bool, error)) *listStream {
	if stream.err != nil || len(stream.list) <= 0 {
		return stream
	}
	var list []interface{}
	var count int
	for i, e := range stream.list {
		//如果flag为true,则保留该元素
		flag, err := filter(i, e)
		if err != nil {
			stream.err = errors.Wrap(err, "Failed at Filter:")
			break
		}
		if flag {
			if list == nil {
				list = make([]interface{}, len(stream.list))
			}
			list = append(list, stream.list[i])
			count++
		}
	}
	//覆盖掉旧的数据
	if count < len(stream.list) {
		stream.list = list[:count]
	}
	return stream
}

//元素排序
func (stream *listStream) Sort(comparator sort.Comparator) *listStream {
	if stream.err != nil || len(stream.list) <= 0 {
		return stream
	}
	//使用sorter进行排序
	if err := sorter.DoSort(stream.list, comparator); err != nil {
		stream.err = errors.Wrap(err,"Failed at Sort:")
	}
	return stream
}

func (stream *listStream) Error() error {
	return stream.err
}
