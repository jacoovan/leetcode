/*
797:所有可能的路径
leetcodeID : 813
leetcode地址 : https://leetcode-cn.com/problems/all-paths-from-source-to-target/
难度 : 中等

给一个有 n 个结点的有向无环图，找到所有从 0 到 n-1 的路径并输出（不要求按顺序）

二维数组的第 i 个数组中的单元都表示有向图中 i 号结点所能到达的下一些结点（译者注：有向图是有方向的，即规定了a&rarr;b你就不能从b&rarr;a）空就是没有下一个结点了。

示例:
输入: [[1,2], [3], [3], []] 
输出: [[0,1,3],[0,2,3]] 
解释: 图是这样的:
0--->1
|    |
v    v
2--->3
这有两条路: 0 -> 1 -> 3 和 0 -> 2 -> 3.


提示:

<ul>
	结点的数量会在范围 [2, 15] 内。
	你可以把路径以任意顺序输出，但在路径内的结点的顺序必须保证。
</ul>

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func allPathsSourceTarget(graph [][]int) [][]int {
    
}