/*
1032:字符流
leetcodeID : 1097
leetcode地址 : https://leetcode-cn.com/problems/stream-of-characters/
难度 : 困难

按下述要求实现 StreamChecker 类：

<ul>
	StreamChecker(words)：构造函数，用给定的字词初始化数据结构。
	query(letter)：如果存在某些 k >= 1，可以用查询的最后 k个字符（按从旧到新顺序，包括刚刚查询的字母）拼写出给定字词表中的某一字词时，返回 true。否则，返回 false。
</ul>

 

示例：

StreamChecker streamChecker = new StreamChecker(["cd","f","kl"]); // 初始化字典
streamChecker.query(&#39;a&#39;);          // 返回 false
streamChecker.query(&#39;b&#39;);          // 返回 false
streamChecker.query(&#39;c&#39;);          // 返回 false
streamChecker.query(&#39;d&#39;);          // 返回 true，因为 &#39;cd&#39; 在字词表中
streamChecker.query(&#39;e&#39;);          // 返回 false
streamChecker.query(&#39;f&#39;);          // 返回 true，因为 &#39;f&#39; 在字词表中
streamChecker.query(&#39;g&#39;);          // 返回 false
streamChecker.query(&#39;h&#39;);          // 返回 false
streamChecker.query(&#39;i&#39;);          // 返回 false
streamChecker.query(&#39;j&#39;);          // 返回 false
streamChecker.query(&#39;k&#39;);          // 返回 false
streamChecker.query(&#39;l&#39;);          // 返回 true，因为 &#39;kl&#39; 在字词表中。

 

提示：

<ul>
	1 <= words.length <= 2000
	1 <= words[i].length <= 2000
	字词只包含小写英文字母。
	待查项只包含小写英文字母。
	待查项最多 40000 个。
</ul>

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

type StreamChecker struct {
    
}


func Constructor(words []string) StreamChecker {
    
}


func (this *StreamChecker) Query(letter byte) bool {
    
}


/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */