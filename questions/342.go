/*
342:4的幂
leetcodeID : 342
leetcode地址 : https://leetcode-cn.com/problems/power-of-four/
难度 : 简单

给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。

示例 1:

输入: 16
输出: true


示例 2:

输入: 5
输出: false

进阶：<br>
你能不使用循环或者递归来完成本题吗？

 */
package main

import (
    "fmt"
    "math"
)

func main(){
    num := math.Pow(4.0, 1.0)


    res := isPowerOfFour(int(num))

    fmt.Println(res)
}

func isPowerOfFour(num int) bool {
    num = 1

    for num != 0 {
        tempNum := num >> 2
        if tempNum << 2 != num {
            return false
        }

        num = tempNum

        if num == 1 {
            return true
        }
    }

    return false
}