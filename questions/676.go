/*
676:实现一个魔法字典
leetcodeID : 676
leetcode地址 : https://leetcode-cn.com/problems/implement-magic-dictionary/
难度 : 中等

实现一个带有buildDict, 以及 search方法的魔法字典。

对于buildDict方法，你将被给定一串不重复的单词来构建一个字典。

对于search方法，你将被给定一个单词，并且判定能否只将这个单词中一个字母换成另一个字母，使得所形成的新单词存在于你构建的字典中。

示例 1:


Input: buildDict(["hello", "leetcode"]), Output: Null
Input: search("hello"), Output: False
Input: search("hhllo"), Output: True
Input: search("hell"), Output: False
Input: search("leetcoded"), Output: False


注意:


	你可以假设所有输入都是小写字母 a-z。
	为了便于竞赛，测试所用的数据量很小。你可以在竞赛结束后，考虑更高效的算法。
	请记住重置MagicDictionary类中声明的类变量，因为静态/类变量会在多个测试用例中保留。 请参阅<a href="http://leetcode.com/faq/#different-output">这里</a>了解更多详情。


 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

type MagicDictionary struct {
    
}


/** Initialize your data structure here. */
func Constructor() MagicDictionary {
    
}


/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string)  {
    
}


/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
    
}


/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */