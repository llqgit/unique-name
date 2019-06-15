package main

import (
	"fmt"
	"math"
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
		r2 = l1 ^ uint32(math.Round(((float64)((1366*r1+150889)%714025)/714025.0)*32767))
		//fmt.Printf("l1: [%b] r1: [%b] l2: [%b] r2: [%b] .. [%b]\n", l1, r1, l2, r2, uint32(math.Round((float64)(((1366*r1+150889)%714025)/714025.0)*12344)))
		l1 = l2
		r1 = r2

	}
	return (r1 << 16) + l1
}

// 打印测试结果
func PrintResult() {
	// 随机取一个 顺序id 区间作为输入
	for i := int(0); i < 11; i++ {
		//num := uint32(math.Pow10(i))
		num := uint32(i)
		name := GetRandomIdName(num)
		fmt.Printf("id: [%v], name: \"%v\"\n", num, name)
	}
}

func main() {
	PrintResult()
}
