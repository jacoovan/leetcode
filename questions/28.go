/*
28:实现 strStr()
leetcodeID : 28
leetcode地址 : https://leetcode-cn.com/problems/implement-strstr/
难度 : 简单

实现 <a href="https://baike.baidu.com/item/strstr/811469" target="_blank">strStr()</a> 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2


示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1


说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 <a href="https://baike.baidu.com/item/strstr/811469" target="_blank">strstr()</a> 以及 Java的 <a href="https://docs.oracle.com/javase/7/docs/api/java/lang/string.html#indexof(java.lang.string)" target="_blank">indexOf()</a> 定义相符。

 */
package main

import(
    "fmt"
    "strings"
)

func main(){
    haystack := "aaaaa"
    needle := "bba"

    res := strStr(haystack, needle)

    fmt.Println(res)
}

func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }

    haystackArr := strings.Split(haystack, "")
    needleArr := strings.Split(needle, "")

    for i, _ := range haystackArr {
        k := i
        over := true

        for _, v1 := range needleArr {
            if v1 != haystackArr[k] {
                over = false
                break
            }
            k++
        }

        if over {
            return i
        }
    }

    return -1
}