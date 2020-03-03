/*
33:搜索旋转排序数组
leetcodeID : 33
leetcode地址 : https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
难度 : 中等

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4


示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1

 */
package main

import(
    "fmt"
)

func main(){
    nums   := []int {4,5,6,7,0,1,2}
    target := 0

    res := search(nums, target)

    fmt.Println(res)

}

func search(nums []int, target int) int {
    low  := 0
    mid  := int((len(nums) - 1) / 2)
    high := len(nums) - 1

    scroll := searchScroll(nums, low, mid, high)

    low1  := 0
    high1 := scroll - 1
    res1 := binary_search(nums, target, low1, high1)
    if res1 > -1 {
        return res1
    }

    low2  := scroll
    high2 := high
    res2 := binary_search(nums, target, low2, high2)
    if res2 > -1 {
        return res2
    }

    return -1
}

func searchScroll(nums []int, low int, mid int, high int) int {
    if high - low <= 2 {
        if nums[low] < nums[high] {
            return low
        } else {
            return high
        }
    }

    if nums[low] > nums[mid] {
        return searchScroll(nums, low, int((low + mid) / 2), mid)
    } else {
        return searchScroll(nums, mid, int((mid + high) / 2), high)
    }

}

func binary_search(nums []int, target int, low int, high int) int {
    if low >= high {
        if nums[low] == target {
            return low
        } else {
            return -1
        }
    }

    if low == high - 1 {
        if nums[low] == target {
            return low
        } else if nums[high] == target {
            return high
        } else {
            return -1
        }
    }

    mid := int((low + high) / 2)

    if nums[low] > target {
        return -1
    } else if nums[low] <= target && nums[mid] >= target {
        return binary_search(nums, target, low, mid)
    } else if nums[mid] <= target && nums[high] >= target {
        return binary_search(nums, target, mid, high)
    } else {
        return -1
    }
}