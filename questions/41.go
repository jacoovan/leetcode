/*
41:缺失的第一个正数
leetcodeID : 41
leetcode地址 : https://leetcode-cn.com/problems/first-missing-positive/
难度 : 困难

给定一个未排序的整数数组，找出其中没有出现的最小的正整数。

示例 1:

输入: [1,2,0]
输出: 3


示例 2:

输入: [3,4,-1,1]
输出: 2


示例 3:

输入: [7,8,9,11,12]
输出: 1


说明:

你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。

 */
package main

import(
    "fmt"
)

func main(){
    nums := []int {3,4,-1,1}
    fmt.Println(nums)
    res := firstMissingPositive(nums)
    fmt.Println(nums, res)
    fmt.Println("")

    nums = []int {0, 2, 1, 3}
    fmt.Println(nums)
    res = firstMissingPositive(nums)
    fmt.Println(nums, res)
    fmt.Println("")

    nums = []int {0, 1, 2}
    fmt.Println(nums)
    res = firstMissingPositive(nums)
    fmt.Println(nums, res)
    fmt.Println("")

    nums = []int {7,8,9,11,12}
    fmt.Println(nums)
    res = firstMissingPositive(nums)
    fmt.Println(nums, res)
    fmt.Println("")
}

func firstMissingPositive(nums []int) int {
    tackleNums(nums)
    res := getFirstNum(nums)
    return res
}

func getFirstNum(nums []int) int {
    if len(nums) == 0 {
        return 1
    }

    if len(nums) == 1 && nums[0] != 1 {
        return 1
    }

    if nums[0] == 1 {
        return 1
    }

    for i, v := range nums {
        if i != v {
            return i
        }
    }

    return len(nums)
}

func tackleNums(nums []int) {
    for i, v := range nums {
        if v > len(nums) - 1 || v < 0 {
            nums[i] = 0
        }
    }

    for i, v := range nums {
        temp := nums[v]
        if temp > 0 {
            if temp < v || v <= 0 {
                nums[v] = v
                nums[i] = temp
            }
        }
    }
}