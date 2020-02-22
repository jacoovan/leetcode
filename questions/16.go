/*
16:最接近的三数之和
leetcodeID : 16
leetcode地址 : https://leetcode-cn.com/problems/3sum-closest/
难度 : 中等

给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).


 */
package main

import(
    "fmt"
    "math"
)

func main(){
    nums := []int {-1, 2, 1, -4}
    target := 1

    res := threeSumClosest(nums, target)

    fmt.Println(res)
}

func threeSumClosest(nums []int, target int) int {
    mininum := -1

    for i := 0; i < len(nums) - 2; i++ {
        for j := i + 1; j < len(nums) - 1; j++ {
            for k := j + 1; k < len(nums); k++ {
                sum := nums[i] + nums[j] + nums[k]
                if mininum == -1 {
                    mininum = sum
                } else {
                    if int(math.Abs(float64(sum - target))) < mininum {
                        mininum = sum
                    }
                }
            }
        }
    }

    return mininum
}