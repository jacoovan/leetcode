/*
9:回文数
leetcodeID : 9
leetcode地址 : https://leetcode-cn.com/problems/palindrome-number/
难度 : 简单

判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true


示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。


示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。


进阶:

你能不将整数转为字符串来解决这个问题吗？

 */
package main

import(
    "fmt"
)

func main(){
    num := -121
    res := isPalindrome(num)
    fmt.Println(res)
}

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    len := 0
    tmp, nums := 0, []int {}
    for {
        if float64(x) / 10 > 0 {
            tmp  = x - int(x / 10) * 10
            nums = append(nums, tmp) 
            x = int(x / 10)
            len++
        } else {
            break
        }
    }

    i, res := 0, true
    for {
        if i > len - 1 {
            break
        }

        if nums[i] != nums[len - 1 - i] {
            res = false
            break
        }

        i++
    }
    return res
}