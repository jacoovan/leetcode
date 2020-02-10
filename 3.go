/*
3:无重复字符的最长子串
leetcodeID : 3
leetcode地址 : https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
难度 : 中等

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。


示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。


示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


 */
package main

import(
    "fmt"
)

func main(){
    s := "pwwkew"
    l := lengthOfLongestSubstring(s)
    fmt.Println(l)
}

func lengthOfLongestSubstring(s string) int {
    runes := []rune(s)
    var current map[string]string
    current = make(map[string]string)

    i, longest := 0, 0
    for {
        if i >= len(runes) - 1 {
            break
        }

        char := string(runes[i])
        if runes[i] == runes[i + 1] {
            current = map[string]string {char:char}
        } else {
            current[char] = char
        }

        if len(current) > longest {
            longest = len(current)
        }

        i++
    }
    return longest
}