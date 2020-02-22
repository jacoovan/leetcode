/*
18:四数之和
leetcodeID : 18
leetcode地址 : https://leetcode-cn.com/problems/4sum/
难度 : 中等

给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]


 */
package main

import(
    "fmt"
    "sort"
    "strconv"
)

func main(){
    nums := []int {1, 0, -1, 0, -2, 2}
    target := 0

    res := fourSum(nums, target)

    fmt.Println(res)
}

func fourSum(nums []int, target int) [][]int {
    lenOfNum := len(nums)

    temp := [][]int {}
    for i := 0; i < lenOfNum - 3; i++ {
        for j := i + 1; j < lenOfNum - 2; j++ {
            for k := j + 1; k < lenOfNum - 1; k++ {
                for l := k + 1; l < lenOfNum; l++ {

                    if nums[i] + nums[j] + nums[k] + nums[l] == target {
                        temp = append(temp, []int {nums[i], nums[j], nums[k], nums[l]})
                    }
                }
            }
        }
    }

    tempMap := map[string][]int {}
    for _, nums := range temp {
        sort.Ints(nums)

        key := ""
        for _, v := range nums {
            key += strconv.Itoa(v)
        }

        tempMap[key] = nums
    }

    res := [][]int {}
    for _, v := range tempMap {
        res = append(res, v)
    }

    return res
}