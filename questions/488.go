/*
488:祖玛游戏
leetcodeID : 488
leetcode地址 : https://leetcode-cn.com/problems/zuma-game/
难度 : 困难

回忆一下祖玛游戏。现在桌上有一串球，颜色有红色(R)，黄色(Y)，蓝色(B)，绿色(G)，还有白色(W)。 现在你手里也有几个球。

每一次，你可以从手里的球选一个，然后把这个球插入到一串球中的某个位置上（包括最左端，最右端）。接着，如果有出现三个或者三个以上颜色相同的球相连的话，就把它们移除掉。重复这一步骤直到桌上所有的球都被移除。

找到插入并可以移除掉桌上所有球所需的最少的球数。如果不能移除桌上所有的球，输出 -1 。


示例:
输入: "WRRBBW", "RB" 
输出: -1 
解释: WRRBBW -> WRR[R]BBW -> WBBW -> WBB[B]W -> WW （翻译者标注：手上球已经用完，桌上还剩两个球无法消除，返回-1）

输入: "WWRRBBWW", "WRBRW" 
输出: 2 
解释: WWRRBBWW -> WWRR[R]BBWW -> WWBBWW -> WWBB[B]WW -> WWWW -> empty

输入:"G", "GGGGG" 
输出: 2 
解释: G -> G[G] -> GG[G] -> empty 

输入: "RBYYBBRRB", "YRBGB" 
输出: 3 
解释: RBYYBBRRB -> RBYY[Y]BBRRB -> RBBBRRB -> RRRB -> B -> B[B] -> BB[B] -> empty 


标注:


	你可以假设桌上一开始的球中，不会有三个及三个以上颜色相同且连着的球。
	桌上的球不会超过20个，输入的数据中代表这些球的字符串的名字是 "board" 。
	你手中的球不会超过5个，输入的数据中代表这些球的字符串的名字是 "hand"。
	输入的两个字符串均为非空字符串，且只包含字符 &#39;R&#39;,&#39;Y&#39;,&#39;B&#39;,&#39;G&#39;,&#39;W&#39;。


 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func findMinStep(board string, hand string) int {
    
}