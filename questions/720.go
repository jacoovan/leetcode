/*
720:词典中最长的单词
leetcodeID : 720
leetcode地址 : https://leetcode-cn.com/problems/longest-word-in-dictionary/
难度 : 简单

给出一个字符串数组words组成的一本英语词典。从中找出最长的一个单词，该单词是由words词典中其他单词逐步添加一个字母组成。若其中有多个可行的答案，则返回答案中字典序最小的单词。

若无答案，则返回空字符串。

示例 1:


输入: 
words = ["w","wo","wor","worl", "world"]
输出: "world"
解释: 
单词"world"可由"w", "wo", "wor", 和 "worl"添加一个字母组成。


示例 2:


输入: 
words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
输出: "apple"
解释: 
"apply"和"apple"都能由词典中的单词组成。但是"apple"得字典序小于"apply"。


注意:

<ul>
	所有输入的字符串都只包含小写字母。
	words数组长度范围为[1,1000]。
	words[i]的长度范围为[1,30]。
</ul>

 */
package main

import(
    "fmt"
    "sort"
)

func main(){
    //words := []string {"w","wo","wor","worl", "world"}
    words := []string {"a", "banana", "app", "appl", "ap", "apply", "apple"}

    res := longestWord(words)

    fmt.Println(res)
}

func longestWord(words []string) string {
    sorts := map[int][]string {}
    temp := map[int]int {}

    for _, word := range words {
        sorts[len(word)] = []string {}
        temp[len(word)] = len(word)
    }

    lengths := []int {}
    for _, v := range temp {
        lengths = append(lengths, v)
    }
    sort.Ints(lengths)

    for _, word := range words {
        sorts[len(word)] = append(sorts[len(word)], word)
    }

    prev := false
    prevWords := []string {}
    prevLength := 0
    for _, length := range lengths {
        if !prev {
            prevWords = sorts[length]
            prev = true
            prevLength = length
            continue
        }

        res := []string {}
        for _, word := range sorts[length] {
            for _, prevWord := range prevWords {
                if word[prevLength - 1] != prevWord[prevLength - 1] {
                    continue
                }

                res = append(res, word)
            }
        }

        if len(res) == 0 {
            break
        }

        prevWords = res
    }

    max := 'z'
    res := ""
    for _, prevWord := range prevWords {
        if rune(prevWord[len(prevWord)-1]) <= max {
            max = rune(prevWord[len(prevWord)-1])
            res = prevWord
        }
    }

    return res
}