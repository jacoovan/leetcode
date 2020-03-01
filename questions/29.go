/*
29:两数相除
leetcodeID : 29
leetcode地址 : https://leetcode-cn.com/problems/divide-two-integers/
难度 : 中等

给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:

输入: dividend = 10, divisor = 3
输出: 3

示例 2:

输入: dividend = 7, divisor = -3
输出: -2

说明:

<ul>
	被除数和除数均为 32 位有符号整数。
	除数不为 0。
	假设我们的环境只能存储 32 位有符号整数，其数值范围是 [&minus;2<sup>31</sup>,  2<sup>31 </sup>&minus; 1]。本题中，如果除法结果溢出，则返回 2<sup>31 </sup>&minus; 1。
</ul>

 */
package main

import (
    "fmt"
    "math"
)

func main(){
    dividend := 7
    divisor := -3

    res:= divide(dividend, divisor)

    fmt.Println(res)
}

func divide(dividend int, divisor int) int {
    res := 0

    beDiv, div := int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor)))

    isLessZero := false
    if dividend + divisor < beDiv + div {
        isLessZero = true
    }

    for beDiv - div > 0 {
        beDiv = beDiv - div
        res++
    }

    if isLessZero {
        res = res * -1
    }

    return res
}