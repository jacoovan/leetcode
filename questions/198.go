/*
198:打家劫舍
leetcodeID : 198
leetcode地址 : https://leetcode-cn.com/problems/house-robber/
难度 : 简单

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

示例 1:

输入: [1,2,3,1]
输出: 4
解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。

示例 2:

输入: [2,7,9,3,1]
输出: 12
解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。


 */
package main

import(
    "fmt"
)

func main(){
    nums := []int {1,2,3,1}

    res := rob(nums)

    fmt.Println(res)
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		if nums[0] < nums[1] {
			return nums[1]
		}
		return nums[0]
	}

    dp := []int {}
    for k, v := range nums {
    	if k == 0 {
    		dp = append(dp, nums[0])
    		continue
    	}
    	if k == 1 {
    		if nums[0] > nums[1] {
    			dp = append(dp, nums[0])
    		} else {
    			dp = append(dp, nums[1])
    		}
    		continue
    	}
    	if dp[k-1] > dp[k-2] + v {
    		dp = append(dp, dp[k-1])
    	} else {
    		dp = append(dp, dp[k-2]+v)
    	}
    }

    return dp[len(nums)-1]
}