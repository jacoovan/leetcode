/*
730:统计不同回文子字符串
leetcodeID : 730
leetcode地址 : https://leetcode-cn.com/problems/count-different-palindromic-subsequences/
难度 : 困难

给定一个字符串 S，找出 S 中不同的非空回文子序列个数，并返回该数字与 10^9 + 7 的模。

通过从 S 中删除 0 个或多个字符来获得子字符序列。

如果一个字符序列与它反转后的字符序列一致，那么它是回文字符序列。

如果对于某个  i，A_i != B_i，那么 A_1, A_2, ... 和 B_1, B_2, ... 这两个字符序列是不同的。

 

示例 1：

输入：
S = &#39;bccb&#39;
输出：6
解释：
6 个不同的非空回文子字符序列分别为：&#39;b&#39;, &#39;c&#39;, &#39;bb&#39;, &#39;cc&#39;, &#39;bcb&#39;, &#39;bccb&#39;。
注意：&#39;bcb&#39; 虽然出现两次但仅计数一次。


示例 2：

输入：
S = &#39;abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba&#39;
输出：104860361
解释：
共有 3104860382 个不同的非空回文子字符序列，对 10^9 + 7 取模为 104860361。


 

提示：

<ul>
	字符串 S 的长度将在[1, 1000]范围内。
	每个字符 S[i] 将会是集合 {&#39;a&#39;, &#39;b&#39;, &#39;c&#39;, &#39;d&#39;} 中的某一个。
</ul>

 

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func countPalindromicSubsequences(S string) int {
    
}