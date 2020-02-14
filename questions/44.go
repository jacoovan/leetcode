/*
44:通配符匹配
leetcodeID : 44
leetcode地址 : https://leetcode-cn.com/problems/wildcard-matching/
难度 : 困难

给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 &#39;?&#39; 和 &#39;*&#39; 的通配符匹配。

&#39;?&#39; 可以匹配任何单个字符。
&#39;*&#39; 可以匹配任意字符串（包括空字符串）。


两个字符串完全匹配才算匹配成功。

说明:

<ul>
	s 可能为空，且只包含从 a-z 的小写字母。
	p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
</ul>

示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。

示例 2:

输入:
s = "aa"
p = "*"
输出: true
解释: &#39;*&#39; 可以匹配任意字符串。


示例 3:

输入:
s = "cb"
p = "?a"
输出: false
解释: &#39;?&#39; 可以匹配 &#39;c&#39;, 但第二个 &#39;a&#39; 无法匹配 &#39;b&#39;。


示例 4:

输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 &#39;*&#39; 可以匹配空字符串, 第二个 &#39;*&#39; 可以匹配字符串 "dce".


示例 5:

输入:
s = "acdcb"
p = "a*c?b"
输入: false

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func isMatch(s string, p string) bool {
    
}