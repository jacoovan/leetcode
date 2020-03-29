/*
371:两整数之和
leetcodeID : 371
leetcode地址 : https://leetcode-cn.com/problems/sum-of-two-integers/
难度 : 简单

不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。

示例 1:

输入: a = 1, b = 2
输出: 3


示例 2:

输入: a = -2, b = 3
输出: 1

 */
package main

import(
    "fmt"
)

func main(){
    a := 3
    b := 6

    res := getSum(a, b)

    fmt.Println(res)
}

func getSum(a int, b int) int {
    for a != 0 {
        a, b = (a & b) << 1, a ^ b
    }

    return b
}