/*
1286:字母组合迭代器
leetcodeID : 1211
leetcode地址 : https://leetcode-cn.com/problems/iterator-for-combination/
难度 : 中等

请你设计一个迭代器类，包括以下内容：

<ul>
	一个构造函数，输入参数包括：一个 有序且字符唯一 的字符串 characters（该字符串只包含小写英文字母）和一个数字 combinationLength 。
	函数 next() ，按 字典序 返回长度为 combinationLength 的下一个字母组合。
	函数 hasNext() ，只有存在长度为 combinationLength 的下一个字母组合时，才返回 True；否则，返回 False。
</ul>

 

示例：

CombinationIterator iterator = new CombinationIterator("abc", 2); // 创建迭代器 iterator

iterator.next(); // 返回 "ab"
iterator.hasNext(); // 返回 true
iterator.next(); // 返回 "ac"
iterator.hasNext(); // 返回 true
iterator.next(); // 返回 "bc"
iterator.hasNext(); // 返回 false


 

提示：

<ul>
	1 <= combinationLength <= characters.length <= 15
	每组测试数据最多包含 10^4 次函数调用。
	题目保证每次调用函数 next 时都存在下一个字母组合。
</ul>

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

type CombinationIterator struct {
    
}


func Constructor(characters string, combinationLength int) CombinationIterator {
    
}


func (this *CombinationIterator) Next() string {
    
}


func (this *CombinationIterator) HasNext() bool {
    
}


/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */