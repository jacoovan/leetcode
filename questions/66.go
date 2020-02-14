/*
66:加一
leetcodeID : 66
leetcode地址 : https://leetcode-cn.com/problems/plus-one/
难度 : 简单

给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。


示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。


 */
package main

import(
    "fmt"
)

func main(){
    nums := []int {1,2,3}

    ret  := plusOne(nums)
    
    fmt.Println(ret)
}

func plusOne(digits []int) []int {
    len := len(digits)
    if len == 0 {
        return []int {}
    }

    add := (bool) (digits[0] == 9)

    var i int
    for i = len - 1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i] = digits[i] + 1
            break
        } else {
            digits[i] = 0
        }
    }

    if add && i == -1 {
        ret := make([]int, len + 1)
        ret[0] = 1
        for i, v := range digits {
            ret[i + 1] = v
        }
        digits = ret
    }

    return digits
}