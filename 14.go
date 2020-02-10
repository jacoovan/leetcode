/*
14:最长公共前缀
leetcodeID : 14
leetcode地址 : https://leetcode-cn.com/problems/longest-common-prefix/
难度 : 简单

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"


示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。


说明:

所有输入只包含小写字母 a-z 。

 */
package main

import(
    "fmt"
    "strings"
)

func main(){
    s := []string {"flower","flow","flight"}
    // s := []string {"dog","racecar","car"}

    res := longestCommonPrefix(s)

    fmt.Println(res)
}

func longestCommonPrefix(strs []string) string {
    res := ""
    prefix := []string {}

    slen := len(strs)

    chars := [][]string {}
    for _, v := range strs {
        chars = append(chars, strings.Split(v, ""))
    }

    end := false
    d := 0
    for {
        if end {
            break
        }

        w := 0

        if d >= len(chars[w]) {
            break
        }

        basic := chars[0][d]
        basic_right := false

        for {
            if end {
                break
            }

            if w > slen - 1 {
                break
            }

            current := chars[w][d]

            if current != basic {
                end = true
                break
            }

            if w == slen - 1 {
                basic_right = true
            }

            w++
        }

        if basic_right {
            prefix = append(prefix, basic)
        }

        d++
    }

    for _, v := range prefix {
        res += v
    }

    return res
}