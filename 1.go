/*
1:两数之和
leetcodeID : 1
leetcode地址 : https://leetcode-cn.com/problems/two-sum/
难度 : 简单

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]


 */
package main

import(
    "fmt"
)

func main(){
    target := twoSum([]int{0,1,2}, 3)

    fmt.Println(target)
}

func twoSum(nums []int, target int) []int {
    var arrmap map[int]int
    arrmap =  make(map[int]int)
    length := len(nums)
    i      := 0
    for {
        if i >= length {
            break
        }
        if nums[i] != target {
            arrmap[i] = nums[i]
        }

        for j, num := range arrmap {
            if i != j && num + nums[i] == target {
                return []int {
                    num,
                    nums[i],
                }
            }
        }
        i++
    }

    return []int {-1, -1}
}