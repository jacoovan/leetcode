/*
994:腐烂的橘子
leetcodeID : 1036
leetcode地址 : https://leetcode-cn.com/problems/rotting-oranges/
难度 : 简单

在给定的网格中，每个单元格可以有以下三个值之一：

<ul>
	值 0 代表空单元格；
	值 1 代表新鲜橘子；
	值 2 代表腐烂的橘子。
</ul>

每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。

返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。

 

示例 1：

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/02/16/oranges.png" style="height: 150px; width: 712px;">

输入：[[2,1,1],[1,1,0],[0,1,1]]
输出：4


示例 2：

输入：[[2,1,1],[0,1,1],[1,0,1]]
输出：-1
解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。


示例 3：

输入：[[0,2]]
输出：0
解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。


 

提示：


	1 <= grid.length <= 10
	1 <= grid[0].length <= 10
	grid[i][j] 仅为 0、1 或 2


 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func orangesRotting(grid [][]int) int {
    
}