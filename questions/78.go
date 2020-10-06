/*
78:子集
leetcodeID : 78
leetcode地址 : https://leetcode-cn.com/problems/subsets/
难度 : 中等

给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

*/
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	res := subsets(nums)
	fmt.Println("res:", res)
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)

	var generate func(pos int, length int, item []int)
	generate = func(pos int, length int, item []int) {

		if len(item) == length {
			tmp := make([]int, len(item))
			copy(tmp, item)
			res = append(res, tmp)
			return
		}
		for i := pos; i < len(nums); i++ {
			item = append(item, nums[i])
			generate(i+1, length, item)
			item = item[:len(item)-1]
		}
	}

	for length := 1; length <= len(nums); length++ {
		item := make([]int, 0, length)
		generate(0, length, item)
	}

	return res
}
