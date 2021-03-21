/*
473:火柴拼正方形
leetcodeID : 473
leetcode地址 : https://leetcode-cn.com/problems/matchsticks-to-square/
难度 : 中等

还记得童话《卖火柴的小女孩》吗？现在，你知道小女孩有多少根火柴，请找出一种能使用所有火柴拼成一个正方形的方法。不能折断火柴，可以把火柴连接起来，并且每根火柴都要用到。

输入为小女孩拥有火柴的数目，每根火柴用其长度表示。输出即为是否能用所有的火柴拼成正方形。

示例 1:


输入: [1,1,2,2,2]
输出: true

解释: 能拼成一个边长为2的正方形，每边两根火柴。


示例 2:


输入: [3,3,3,3,4]
输出: false

解释: 不能用所有火柴拼成一个正方形。


注意:


	给定的火柴长度和在 0 到 10^9之间。
	火柴数组的长度不超过15。


*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{1, 1, 2, 2, 2}
	nums := []int{3, 3, 3, 3, 4}
	res := makesquare(nums)
	fmt.Println("res:", res)
}

func makesquare(nums []int) bool {
	totalNum := 0
	for _, num := range nums {
		totalNum += num
	}

	tempSideLength := float64(totalNum) / float64(4)
	sideLength := int(tempSideLength)
	for _, num := range nums {
		if num > sideLength {
			return false
		}
	}

	sortNums := sort.IntSlice(nums)
	alreadyUsedMap := make(map[int]bool)
	for i := 0; i < 4; i++ {
		indexes := getSide(sortNums, sideLength, alreadyUsedMap)
		if len(indexes) == 0 {
			return false
		}
		for _, index := range indexes {
			alreadyUsedMap[index] = true
		}
	}
	return true
}

func getSide(nums []int, sideLength int, alreadyUsedMap map[int]bool) []int {
	indexes := make([]int, 0)

	var tempSide int
	for i, smaller := range nums {
		_, ok := alreadyUsedMap[i]
		if ok {
			continue
		}
		tempSide += smaller
		indexes = append(indexes, i)
		if tempSide == sideLength {
			return indexes
		}

		for j := len(nums) - 1; j > i; j-- {
			larger := nums[j]
			if tempSide+larger == sideLength {
				tempSide += larger
				indexes = append(indexes, j)
				return indexes
			}
			if tempSide+larger < sideLength {
				tempSide += larger
				indexes = append(indexes, j)
			}
		}
	}
	if tempSide != sideLength {
		return []int{}
	}
	return indexes
}
