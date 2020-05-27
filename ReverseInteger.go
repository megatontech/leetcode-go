package main
import (
	"fmt"
	"math"
	"strconv"
	"strings"
)
//7. Reverse Integer
//Easy
//Given a 32-bit signed integer, reverse digits of an integer.
//
//Example 1:
//Input: 123
//Output: 321
//
//Example 2:
//Input: -123
//Output: -321
//
//Example 3:
//Input: 120
//Output: 21
//Note:
//Assume we are dealing with an environment which could only store integers
//within the 32-bit signed integer range: [−231,  231 − 1].
//	For the purpose of this problem,
//assume that your function returns 0 when the reversed integer overflows.
func reverse(x int) int {
	return reverse1(x)
}
func reverse1(x int) int {
	//reverse1
	//整数转字符串
	var str = strconv.itoa(x)
	////字符串转整型
	bytes := []rune(str)
	for from , to := 0 , len(bytes) -1 ; from < to ; from , to = from + 1, to -1{
		bytes[from] , bytes[to] = bytes[to] , bytes[from]
	}
	str = string(bytes)
	var result = strconv.Atoi(str)
	return result
}
func reverse2(x int) int {
	//切割为数组，反转后按幂乘然后相加
	//var arr = []int
	var reversearr = []int
	//strings.Split(strconv.Itoa(x))
	bytes := []rune(strconv.Itoa(x))
	for from , to := 0 , len(bytes) -1 ; from < to ; from , to = from + 1, to -1{
		bytes[from] , bytes[to] = bytes[to] , bytes[from]
	}
	for i, i2 := range bytes {
		reversearr[i] = strconv.Atoi(i2)
	}
	var result = 0
	for i, i2 := range reversearr {
		result+=i2*math.Pow10(i)
	}
	result = (int)result
	return  result
}