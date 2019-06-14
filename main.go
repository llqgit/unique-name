package main

import (
	"fmt"
)

// 获取根据 index 打乱顺序的 数字名称（uint 防止首位变成负数）
func GetRandomIdName(index uint32) uint32 {
	l1 := (index >> 16) & 0xFFFF
	r1 := index & 0xFFFF
	var l2 uint32
	var r2 uint32

	// 循环 4 轮
	for i := 0; i < 4; i++ {

		l2 = r1
		r2 = l1 ^ 123456789
		//r2 = l1 ^ rand.Int31n(int32(float32(((1366*r1+150889)%714025)/714025.0)*32767))
		l1 = l2
		r1 = r2
		//fmt.Printf("l1: [%v], r1: [%v], l2: [%v], r2: [%v]\n", l1, r1, l2, r2)
	}
	return (r1 << 16) + l1
}

// 打印测试结果
func PrintResult() {
	// 随机取一个 顺序id 区间作为输入
	for i := uint32(10000); i < 10020; i++ {
		name := GetRandomIdName(i)
		fmt.Printf("id: [%v], name: \"%v\"\n", i, name)
	}
}

func main() {
	PrintResult()
}
