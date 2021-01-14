/*
856:括号的分数
leetcodeID : 886
leetcode地址 : https://leetcode-cn.com/problems/score-of-parentheses/
难度 : 中等

给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：

<ul>
	() 得 1 分。
	AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
	(A) 得 2 * A 分，其中 A 是平衡括号字符串。
</ul>



示例 1：

输入： "()"
输出： 1


示例 2：

输入： "(())"
输出： 2


示例 3：

输入： "()()"
输出： 2


示例 4：

输入： "(()(()))"
输出： 6




提示：


	S 是平衡括号字符串，且只含有 ( 和 ) 。
	2 <= S.length <= 50


*/
package main

import (
	"fmt"
)

func main() {
	s := `(()(()))`
	res := scoreOfParentheses(s)
	fmt.Println("res:", res)
}

func scoreOfParentheses(S string) int {
	return scoreOfSub(S)
}

func scoreOfSub(s string) int {
	score := 0
	if len(s) < 2 {
		return 0
	}
	for i, v := range s[:len(s)-1] {
		if v == '(' && s[i+1] == '(' {
			score = score
			score = score + 2*scoreOfSub(s[i+1:])
			return score
		}
		if v == '(' && s[i+1] == ')' {
			score = score + 1
			score = score + scoreOfSub(s[i+1:])
			return score
		}
		if v == ')' && s[i+1] == '(' {
			score = score
			score = score + scoreOfSub(s[i+1:])
			return score
		}
		if v == ')' && s[i+1] == ')' {
			score = score + 0
			score = score + scoreOfSub(s[i+1:])
			return score
		}
	}
	return score
}
