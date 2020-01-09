package common

import "errors"

// 声明新切片类型
type units []uint32

// 返回切片长度
func (x units) Len() int {
	return len(x)
}

// 对比两个数的大小
func (x units) Less(i, j int) bool {
	return x[i] < x[j]
}

// 当hash环上没有数据时, 提示错误
var errEmpty = errors.New("Hash 环没有数据")
