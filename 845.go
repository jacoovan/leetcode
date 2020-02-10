/*
845:数组中的最长山脉
leetcodeID : 875
leetcode地址 : https://leetcode-cn.com/problems/longest-mountain-in-array/
难度 : 中等

我们把数组 A 中符合下列属性的任意连续子数组 B 称为 &ldquo;山脉&rdquo;：

<ul>
	B.length >= 3
	存在 0 < i < B.length - 1 使得 B[0] < B[1] < ... B[i-1] < B[i] > B[i+1] > ... > B[B.length - 1]
</ul>

（注意：B 可以是 A 的任意子数组，包括整个数组 A。）

给出一个整数数组 A，返回最长 &ldquo;山脉&rdquo; 的长度。

如果不含有 &ldquo;山脉&rdquo; 则返回 0。

 

示例 1：

输入：[2,1,4,7,3,2,5]
输出：5
解释：最长的 &ldquo;山脉&rdquo; 是 [1,4,7,3,2]，长度为 5。


示例 2：

输入：[2,2,2]
输出：0
解释：不含 &ldquo;山脉&rdquo;。


 

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

func longestMountain(A []int) int {
    
}