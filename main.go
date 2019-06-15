package main

import (
	"fmt"
	"math"
)

// 通过10进制的位数，获取对应2进制需要的最大位数
func GetMaxBinaryDigitFromDecimalDigit(length int) uint {
	dec := math.Pow10(length) // 10 进制位数最大值
	count := uint(0)          // 二进制 位数
	for i := 1.0; math.Pow(2, i) < dec; i++ {
		count++
	}
	return count
}

// 获取根据 index 打乱顺序的 数字名称（uint 防止首位变成负数） 最大长度为 length
// num 的位数一定要小于 length，使得 2^length 能够包含 num，防止生成的结果位数超过 length
func GetRandomIdName(num uint64, length int) uint64 {

	// 根据此二进制位数，将数字拆成长度相等的两个部分
	bits := uint(math.Floor(float64(GetMaxBinaryDigitFromDecimalDigit(length)) / 2))
	//fmt.Printf("digit: [%v]\n", bits)

	l1 := num >> bits             // 头部
	r1 := num & ((1 << bits) - 1) // 尾部
	var l2, r2 uint64

	// 循环 3 轮
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ (uint64(math.Round(((float64)((1366*r1+150889)%714025)/714025.0)*32767*32767)))&((1<<bits)-1)
		//fmt.Printf("l1: [%b] r1: [%b] l2: [%b] r2: [%b] .. [%b]\n", l1, r1, l2, r2, uint64(math.Round((float64)(((1366*r1+150889)%714025)/714025.0)*12344)))
		l1 = l2
		r1 = r2

	}
	result := (r1 << bits) + l1
	return result
}

// 打印测试结果
func PrintResult() {
	// 随机取一个 顺序id 区间作为输入
	for i := int(0); i < 20; i++ {
		//num := uint64(math.Pow10(i))
		num := uint64(i)
		name := GetRandomIdName(num, 5)
		fmt.Printf("id: [%v], name: \"%v\"\n", num, name)
	}
}

func main() {
	PrintResult()
}
