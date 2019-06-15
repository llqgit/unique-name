package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// 检测两次计算的结果是否相同
func CheckUniqueResult(list []uint64) bool {
	for i := 0; i < len(list); i++ {
		//num := uint64(math.Pow10(i))
		numA := uint64(list[i])
		numB := uint64(list[i])
		nameA := GetRandomIdName(numA, 5)
		nameB := GetRandomIdName(numB, 5)
		if nameA != nameB {
			return false
		}
	}
	return true
}

// 检查是否有重复项
func CheckRepeat(list []uint64) bool {
	tempMap := make(map[uint64]bool)
	for i := 0; i < len(list); i++ {
		num := uint64(list[i])
		// 获取随机的数字名称
		name := GetRandomIdName(num, 5)
		// 如果已经存在重复的 name 则返回 false
		if hasValue, ok := tempMap[name]; ok && hasValue {
			return false
		}
		// 将新的记录压入 map
		tempMap[name] = true
	}
	return true
}

// 检测生成的值长度是否有超出最大长度的
func CheckMaxLength(list []uint64, length int) bool {
	logLenCount := make(map[int]int)
	totalCalcCount := 0
	isSuccess := true
	for i := 0; i < len(list); i++ {
		num := uint64(list[i])
		name := GetRandomIdName(num, length)
		str := fmt.Sprint(name)
		// 记录长度个数
		if _, ok := logLenCount[len(str)]; ok {
			logLenCount[len(str)]++
		} else {
			logLenCount[len(str)] = 1
		}
		totalCalcCount++
		if len(str) > length {
			isSuccess = false
		}
	}
	//fmt.Printf("str: [%v] len: [%v] logCount: [%v]\n", str, len(str), logLenCount)
	fmt.Printf(" totalCount: [%v] remainCount: [%v]\n logLenCount: %v", totalCalcCount, len(list)-totalCalcCount, logLenCount)
	return isSuccess
}

// 检查进制转换长度方法
func CheckDigitCalc() bool {
	// 10 2^3 8
	if digit := GetMaxBinaryDigitFromDecimalDigit(1); digit != 3 {
		fmt.Printf("digit [%v] should be 3", digit)
		return false
	}
	// 100 2^6 64
	if digit := GetMaxBinaryDigitFromDecimalDigit(2); digit != 6 {
		fmt.Printf("digit [%v] should be 6", digit)
		return false
	}
	// 1000 2^9 512
	if digit := GetMaxBinaryDigitFromDecimalDigit(3); digit != 9 {
		fmt.Printf("digit [%v] should be 9", digit)
		return false
	}
	// 10000 2^13 8192
	if digit := GetMaxBinaryDigitFromDecimalDigit(4); digit != 13 {
		fmt.Printf("digit [%v] should be 13", digit)
		return false
	}
	return true
}

// 测试结果唯一性
func TestUniqueResult(t *testing.T) {
	// 检查多次计算是否结果唯一
	Convey("TestUniqueResult should return true from [0] to [10,000]", t, func() {
		a := uint64(0)
		b := uint64(10000)
		list := make([]uint64, 0)
		for i := a; i < b; i++ {
			list = append(list, i)
		}
		So(CheckUniqueResult(list), ShouldBeTrue)
	})
}

// 测试重复性
func TestRepeat(t *testing.T) {
	// 测试一定有重复
	Convey("Test has repeat should return false", t, func() {
		list := make([]uint64, 0)
		list = append(list, 0, 0)
		So(CheckRepeat(list), ShouldBeFalse)
	})
	// 检查是否有重复
	Convey("TestRepeat should return true from [0] to [1,000,000]", t, func() {
		a := uint64(0)
		b := uint64(1000000)
		list := make([]uint64, 0)
		for i := a; i < b; i++ {
			list = append(list, i)
		}
		So(CheckRepeat(list), ShouldBeTrue)
	})
}

// 测试进制转换长度方法
func TestDigitLengthFunc(t *testing.T) {
	// 检查是否有超出最大长度的结果值
	Convey("TestDigitLengthFunc should return true", t, func() {
		So(CheckDigitCalc(), ShouldBeTrue)
	})
}

// 测试最大长度
func TestLength(t *testing.T) {
	// 检查是否有超出最大长度的结果值
	Convey("TestLength should return true from [0] to [1,000,000]", t, func() {
		a := uint64(0)
		b := uint64(100000)
		list := make([]uint64, 0)
		for i := a; i < b; i++ {
			list = append(list, i)
		}
		So(CheckMaxLength(list, 6), ShouldBeTrue)
	})
}
