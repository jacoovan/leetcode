/*
10:正则表达式匹配
leetcodeID : 10
leetcode地址 : https://leetcode-cn.com/problems/regular-expression-matching/
难度 : 困难

给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 &#39;.&#39; 和 &#39;*&#39; 的正则表达式匹配。

&#39;.&#39; 匹配任意单个字符
&#39;*&#39; 匹配零个或多个前面的那一个元素


所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

<ul>
	s 可能为空，且只包含从 a-z 的小写字母。
	p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
</ul>

示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。


示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 &#39;*&#39; 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 &#39;a&#39;。因此，字符串 "aa" 可被视为 &#39;a&#39; 重复了一次。


示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（&#39;*&#39;）任意字符（&#39;.&#39;）。


示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 &#39;*&#39; 表示零个或多个，这里 &#39;c&#39; 为 0 个, &#39;a&#39; 被重复一次。因此可以匹配字符串 "aab"。


示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false

 */
package main

import(
    "fmt"
)

func main(){
    s := "aa"
    p := "a*."
    ret := isMatch(s, p)
    fmt.Println(ret)
}

func isMatch(s string, p string) bool {
    sRunes := []rune(s)
    pRunes := []rune(p)

    pIndex := 0
    for _, sCurrent := range sRunes {
        pCurrent := pRunes[pIndex]

        if pCurrent == 46 {
            pCurrent = pRunes[pIndex - 1]
        } else if pCurrent == 42 {
            if sCurrent == pRunes[pIndex - 1] {
                pIndex--
            } else {
                pIndex++
            }
            pCurrent = pRunes[pIndex]
        }

        if sCurrent != pCurrent {
            return false
        }
        pIndex++
    }

    for {
        if pIndex <= len(pRunes) - 1 {
            if pRunes[pIndex] != 42 {
                return false
            }
            pIndex++
        } else if pIndex > len(pRunes) - 1 {
            break
        }
    }
    

    if pIndex < len(pRunes) - 1 {
        return false
    }

    return true
}