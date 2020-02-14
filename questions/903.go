/*
903:DI 序列的有效排列
leetcodeID : 939
leetcode地址 : https://leetcode-cn.com/problems/valid-permutations-for-di-sequence/
难度 : 困难

我们给出 S，一个源于 {&#39;D&#39;, &#39;I&#39;} 的长度为 n 的字符串 。（这些字母代表 &ldquo;减少&rdquo; 和 &ldquo;增加&rdquo;。）<br>
有效排列 是对整数 {0, 1, ..., n} 的一个排列 P[0], P[1], ..., P[n]，使得对所有的 i：

<ul>
	如果 S[i] == &#39;D&#39;，那么 P[i] > P[i+1]，以及；
	如果 S[i] == &#39;I&#39;，那么 P[i] < P[i+1]。
</ul>

有多少个有效排列？因为答案可能很大，所以请返回你的答案模 10^9 + 7.

 

示例：

输入："DID"
输出：5
解释：
(0, 1, 2, 3) 的五个有效排列是：
(1, 0, 3, 2)
(2, 0, 3, 1)
(2, 1, 3, 0)
(3, 0, 2, 1)
(3, 1, 2, 0)


 

提示：


	1 <= S.length <= 200
	S 仅由集合 {&#39;D&#39;, &#39;I&#39;} 中的字符组成。


 

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func numPermsDISequence(S string) int {
    
}