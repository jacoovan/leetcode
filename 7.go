/*
7:整数反转
leetcodeID : 7
leetcode地址 : https://leetcode-cn.com/problems/reverse-integer/
难度 : 简单

给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321


 示例 2:

输入: -123
输出: -321


示例 3:

输入: 120
输出: 21


注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [&minus;2<sup>31</sup>,  2<sup>31 </sup>&minus; 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

 */
package main

import(
    "fmt"
    "math"
)

func main(){
    n := -123
    n1:= reverse(n)
    fmt.Println(n1)
}

func reverse(x int) int {
    if x == -2 ^ 31 {
        return 0
    }

    nums := []int {}
    nums = r1(x, nums)
    res := 0
    if len(nums) > 0 {
        for i := 0; i <= len(nums) - 1; i++ {
            res += nums[i] * int(math.Pow(10, float64(len(nums) - 1 - i)))
        }
    }
    return res
}

func r1(x int, nums []int) []int{
    if math.Abs(float64(x)) < 1 {
        return nums
    }

    nums = append(nums, x % 10)
    x    = int((x - x % 10) / 10)

    nums = r1(x, nums)

    return nums
}