/*
1138:字母板上的路径
leetcodeID : 1238
leetcode地址 : https://leetcode-cn.com/problems/alphabet-board-path/
难度 : 中等

我们从一块字母板上的位置 (0, 0) 出发，该坐标对应的字符为 board[0][0]。

在本题里，字母板为board = ["abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"].

我们可以按下面的指令规则行动：

<ul>
	如果方格存在，&#39;U&#39; 意味着将我们的位置上移一行；
	如果方格存在，&#39;D&#39; 意味着将我们的位置下移一行；
	如果方格存在，&#39;L&#39; 意味着将我们的位置左移一列；
	如果方格存在，&#39;R&#39; 意味着将我们的位置右移一列；
	&#39;!&#39; 会把在我们当前位置 (r, c) 的字符 board[r][c] 添加到答案中。
</ul>

返回指令序列，用最小的行动次数让答案和目标 target 相同。你可以返回任何达成目标的路径。

 

示例 1：

输入：target = "leet"
输出："DDR!UURRR!!DDD!"


示例 2：

输入：target = "code"
输出："RR!DDRR!UUL!R!"


 

提示：

<ul>
	1 <= target.length <= 100
	target 仅含有小写英文字母。
</ul>

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func alphabetBoardPath(target string) string {
    
}