/*
15:三数之和
leetcodeID : 15
leetcode地址 : https://leetcode-cn.com/problems/3sum/
难度 : 中等

给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]


 */
package main

import(
    "fmt"
    "sort"
    "strconv"
)

func main(){
    nums := []int {-1, 0, 1, 2, -1, -4}

    res := threeSum(nums)

    fmt.Println(res)
}

func threeSum(nums []int) [][]int {
    res, temp := [][]int {}, [][]int {}

    // 取出数组中和为0的不同位置数字
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                if nums[i] + nums[j] + nums[k] == 0 {
                    temp = append(temp, []int {nums[i], nums[j], nums[k]})
                }
            }
        }
    }

    // 对相同数字组合进行去重
    temp_map := map[string][]int {}
    for _, v := range temp {
        sort.Ints(v)
        key := ""
        for _, num := range v {
            key += strconv.Itoa(num)
        }
        temp_map[key] = v
    }

    // 得到最终结果
    for _, v := range temp_map {
        res = append(res, v)
    }

    return res
}
