/*
1041:困于环中的机器人
leetcodeID : 1119
leetcode地址 : https://leetcode-cn.com/problems/robot-bounded-in-circle/
难度 : 中等

在无限的平面上，机器人最初位于 (0, 0) 处，面朝北方。机器人可以接受下列三条指令之一：

<ul>
	"G"：直走 1 个单位
	"L"：左转 90 度
	"R"：右转 90 度
</ul>

机器人按顺序执行指令 instructions，并一直重复它们。

只有在平面中存在环使得机器人永远无法离开时，返回 true。否则，返回 false。

 

示例 1：

输入："GGLLGG"
输出：true
解释：
机器人从 (0,0) 移动到 (0,2)，转 180 度，然后回到 (0,0)。
重复这些指令，机器人将保持在以原点为中心，2 为半径的环中进行移动。


示例 2：

输入："GG"
输出：false
解释：
机器人无限向北移动。


示例 3：

输入："GL"
输出：true
解释：
机器人按 (0, 0) -> (0, 1) -> (-1, 1) -> (-1, 0) -> (0, 0) -> ... 进行移动。

 

提示：


	1 <= instructions.length <= 100
	instructions[i] 在 {&#39;G&#39;, &#39;L&#39;, &#39;R&#39;} 中


 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func isRobotBounded(instructions string) bool {
    
}