/*
207:课程表
leetcodeID : 207
leetcode地址 : https://leetcode-cn.com/problems/course-schedule/
难度 : 中等

现在你总共有 n 门课需要选，记为 0 到 n-1。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]

给定课程总量以及它们的先决条件，判断是否可能完成所有课程的学习？

示例 1:

输入: 2, [[1,0]] 
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。

示例 2:

输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。

说明:


	输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见<a href="http://blog.csdn.net/woaidapaopao/article/details/51732947" target="_blank">图的表示法</a>。
	你可以假定输入的先决条件中没有重复的边。


提示:


	这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
	<a href="https://www.coursera.org/specializations/algorithms" target="_blank">通过 DFS 进行拓扑排序</a> - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
	
	拓扑排序也可以通过 <a href="https://baike.baidu.com/item/%e5%ae%bd%e5%ba%a6%e4%bc%98%e5%85%88%e6%90%9c%e7%b4%a2/5224802?fr=aladdin&fromid=2148012&fromtitle=%e5%b9%bf%e5%ba%a6%e4%bc%98%e5%85%88%e6%90%9c%e7%b4%a2" target="_blank">BFS</a> 完成。
	


 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func canFinish(numCourses int, prerequisites [][]int) bool {
    
}