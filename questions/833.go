/*
833:字符串中的查找与替换
leetcodeID : 862
leetcode地址 : https://leetcode-cn.com/problems/find-and-replace-in-string/
难度 : 中等

对于某些字符串 S，我们将执行一些替换操作，用新的字母组替换原有的字母组（不一定大小相同）。

每个替换操作具有 3 个参数：起始索引 i，源字 x 和目标字 y。规则是如果 x 从原始字符串 S 中的位置 i 开始，那么我们将用 y 替换出现的 x。如果没有，我们什么都不做。

举个例子，如果我们有 S = &ldquo;abcd&rdquo; 并且我们有一些替换操作 i = 2，x = &ldquo;cd&rdquo;，y = &ldquo;ffff&rdquo;，那么因为 &ldquo;cd&rdquo; 从原始字符串 S 中的位置 2 开始，我们将用 &ldquo;ffff&rdquo; 替换它。

再来看 S = &ldquo;abcd&rdquo; 上的另一个例子，如果我们有替换操作 i = 0，x = &ldquo;ab&rdquo;，y = &ldquo;eee&rdquo;，以及另一个替换操作 i = 2，x = &ldquo;ec&rdquo;，y = &ldquo;ffff&rdquo;，那么第二个操作将不执行任何操作，因为原始字符串中 S[2] = &#39;c&#39;，与 x[0] = &#39;e&#39; 不匹配。

所有这些操作同时发生。保证在替换时不会有任何重叠： S = "abc", indexes = [0, 1], sources = ["ab","bc"] 不是有效的测试用例。

 

示例 1：

输入：S = "abcd", indexes = [0,2], sources = ["a","cd"], targets = ["eee","ffff"]
输出："eeebffff"
解释：
"a" 从 S 中的索引 0 开始，所以它被替换为 "eee"。
"cd" 从 S 中的索引 2 开始，所以它被替换为 "ffff"。


示例 2：

输入：S = "abcd", indexes = [0,2], sources = ["ab","ec"], targets = ["eee","ffff"]
输出："eeecd"
解释：
"ab" 从 S 中的索引 0 开始，所以它被替换为 "eee"。
"ec" 没有从原始的 S 中的索引 2 开始，所以它没有被替换。


 

提示：


	0 <= indexes.length = sources.length = targets.length <= 100
	0 < indexes[i] < S.length <= 1000
	给定输入中的所有字符都是小写字母。


 

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
    
}