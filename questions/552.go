/*
552:学生出勤记录 II
leetcodeID : 552
leetcode地址 : https://leetcode-cn.com/problems/student-attendance-record-ii/
难度 : 困难

给定一个正整数 n，返回长度为 n 的所有可被视为可奖励的出勤记录的数量。 答案可能非常大，你只需返回结果mod 10<sup>9</sup> + 7的值。

学生出勤记录是只包含以下三个字符的字符串：


	&#39;A&#39; : Absent，缺勤
	&#39;L&#39; : Late，迟到
	&#39;P&#39; : Present，到场


如果记录不包含多于一个&#39;A&#39;（缺勤）或超过两个连续的&#39;L&#39;（迟到），则该记录被视为可奖励的。

示例 1:


输入: n = 2
输出: 8 
解释：
有8个长度为2的记录将被视为可奖励：
"PP" , "AP", "PA", "LP", "PL", "AL", "LA", "LL"
只有"AA"不会被视为可奖励，因为缺勤次数超过一次。

注意：n 的值不会超过100000。

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func checkRecord(n int) int {
    
}