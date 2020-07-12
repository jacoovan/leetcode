/*
746:使用最小花费爬楼梯
leetcodeID : 747
leetcode地址 : https://leetcode-cn.com/problems/min-cost-climbing-stairs/
难度 : 简单

数组的每个索引做为一个阶梯，第 i个阶梯对应着一个非负数的体力花费值 cost[i](索引从0开始)。

每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。

您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。

示例 1:


输入: cost = [10, 15, 20]
输出: 15
解释: 最低花费是从cost[1]开始，然后走两步即可到阶梯顶，一共花费15。


 示例 2:


输入: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出: 6
解释: 最低花费方式是从cost[0]开始，逐个经过那些1，跳过cost[3]，一共花费6。


注意：


	cost 的长度将会在 [2, 1000]。
	每一个 cost[i] 将会是一个Integer类型，范围为 [0, 999]。


 */
package main

import(
    "fmt"
)

func main(){
    cost := []int{0, 0, 1, 1}
    res := minCostClimbingStairs(cost)
    fmt.Println(res)
}

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	if len(cost) == 3 {
		return min(cost[1], cost[0]+cost[2])
	}


	return min(
		cost[len(cost)-1] + minCostClimbingStairs(cost[:len(cost)-1]),
		cost[len(cost)-2] + minCostClimbingStairs(cost[:len(cost)-2]),
	)
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
