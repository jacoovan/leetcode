/*
1310:子数组异或查询
leetcodeID : 1435
leetcode地址 : https://leetcode-cn.com/problems/xor-queries-of-a-subarray/
难度 : 中等

有一个正整数数组 arr，现给你一个对应的查询数组 queries，其中 queries[i] = [L<sub>i, </sub>R<sub>i</sub>]。

对于每个查询 i，请你计算从 L<sub>i</sub> 到 R<sub>i</sub> 的 XOR 值（即 arr[L<sub>i</sub>] xor arr[L<sub>i+1</sub>] xor ... xor arr[R<sub>i</sub>]）作为本次查询的结果。

并返回一个包含给定查询 queries 所有结果的数组。

 

示例 1：

输入：arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]
输出：[2,7,14,8] 
解释：
数组中元素的二进制表示形式是：
1 = 0001 
3 = 0011 
4 = 0100 
8 = 1000 
查询的 XOR 值为：
[0,1] = 1 xor 3 = 2 
[1,2] = 3 xor 4 = 7 
[0,3] = 1 xor 3 xor 4 xor 8 = 14 
[3,3] = 8


示例 2：

输入：arr = [4,8,2,10], queries = [[2,3],[1,3],[0,0],[0,3]]
输出：[8,0,4,4]


 

提示：

<ul>
	1 <= arr.length <= 3 * 10^4
	1 <= arr[i] <= 10^9
	1 <= queries.length <= 3 * 10^4
	queries[i].length == 2
	0 <= queries[i][0] <= queries[i][1] < arr.length
</ul>

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func xorQueries(arr []int, queries [][]int) []int {
    
}