/*
212:单词搜索 II
leetcodeID : 212
leetcode地址 : https://leetcode-cn.com/problems/word-search-ii/
难度 : 困难

给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中&ldquo;相邻&rdquo;单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

示例:

输入: 
words = ["oath","pea","eat","rain"] and board =
[
  [&#39;o&#39;,&#39;a&#39;,&#39;a&#39;,&#39;n&#39;],
  [&#39;e&#39;,&#39;t&#39;,&#39;a&#39;,&#39;e&#39;],
  [&#39;i&#39;,&#39;h&#39;,&#39;k&#39;,&#39;r&#39;],
  [&#39;i&#39;,&#39;f&#39;,&#39;l&#39;,&#39;v&#39;]
]

输出: ["eat","oath"]

说明:<br>
你可以假设所有输入都由小写字母 a-z 组成。

提示:

<ul>
	你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
	如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？ 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： <a href="/problems/implement-trie-prefix-tree/description/">实现Trie（前缀树）</a>。
</ul>

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func findWords(board [][]byte, words []string) []string {
    
}