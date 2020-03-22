/*
231:2的幂
leetcodeID : 231
leetcode地址 : https://leetcode-cn.com/problems/power-of-two/
难度 : 简单

给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

示例 1:

输入: 1
输出: true
解释: 2<sup>0</sup> = 1

示例 2:

输入: 16
输出: true
解释: 2<sup>4</sup> = 16

示例 3:

输入: 218
输出: false

 */
package main

import "fmt"

func main(){
    num := 4

    res := isPowerOfTwo(num)

    fmt.Println(res)
}

func isPowerOfTwo(n int) bool {
    count := 0
    for n != 0 {
        count++
        n &= n - 1
    }

    return count == 1
}