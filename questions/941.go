/*
941:有效的山脉数组
leetcodeID : 978
leetcode地址 : https://leetcode-cn.com/problems/valid-mountain-array/
难度 : 简单

给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。

让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：

<ul>
	A.length >= 3
	在 0 < i < A.length - 1 条件下，存在 i 使得：
	<ul>
		A[0] < A[1] < ... A[i-1] < A[i] 
		A[i] > A[i+1] > ... > A[B.length - 1]
	</ul>
	
</ul>

 

示例 1：

输入：[2,1]
输出：false


示例 2：

输入：[3,5,5]
输出：false


示例 3：

输入：[0,3,2,1]
输出：true

 

提示：


	0 <= A.length <= 10000
	0 <= A[i] <= 10000 


 

 

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func validMountainArray(A []int) bool {
    
}