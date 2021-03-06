/*
34:在排序数组中查找元素的第一个和最后一个位置
leetcodeID : 34
leetcode地址 : https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
难度 : 中等

给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]

示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

 */
package main

import "fmt"

func main(){
    nums := []int {5,7,7,8,8,10}
    target := 6
    res := searchRange(nums, target)

    fmt.Println(res)
}

func searchRange(nums []int, target int) []int {
    res := []int {-1, -1}
    first := true

    for i, v := range nums {
        if v == target {
            if first {
                res[0] = i
                first = false
            }

            res[1] = i
        }
    }

    return res
}