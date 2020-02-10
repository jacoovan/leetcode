/*
5:最长回文子串
leetcodeID : 5
leetcode地址 : https://leetcode-cn.com/problems/longest-palindromic-substring/
难度 : 中等

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。


示例 2：

输入: "cbbd"
输出: "bb"


 */
package main

import(
    "fmt"
)

func main(){
    s := "babad"
    s1:= longestPalindrome(s)
    fmt.Println(s1)
}

func longestPalindrome(s string) string {
    bytes := []rune(s)

    len := len(bytes)
    max := 0
    start, end := -1, -1
    for i := 0; i <= len - 1; i++{
        for j := i; j <= len - 1; j++ {
            ret := isPalindrome(bytes[i:j+1])
            current := j - i + 1
            if ret && current > max {
                max = current
                start = i
                end   = j + 1
            }
        }
    }

    if start == -1 {
        return ""
    }
    return string(bytes[start:end])
}

func isPalindrome(s []rune) bool {
    len := len(s)

    if len <= 1 {
        return true
    }

    i := 0
    j := len - 1
    for {
        if i >= j {
            break
        }
        if s[i] != s[j]{
            break
        }

        i++
        j--
    }

    if len % 2 == 1 && i == j {
        return true
    } else if i == j + 1 {
        return true
    }

    return false
}