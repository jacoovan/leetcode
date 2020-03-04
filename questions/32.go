/*
32:最长有效括号
leetcodeID : 32
leetcode地址 : https://leetcode-cn.com/problems/longest-valid-parentheses/
难度 : 困难

给定一个只包含 &#39;(&#39; 和 &#39;)&#39; 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"


示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"


 */
package main

import(
    "fmt"
    "strings"
)

func main(){
    s := ")()())"
    res := longestValidParentheses(s)
    fmt.Println(res)
}

func longestValidParentheses(s string) int {
    temp := map[int]map[int][]string {}
    chars := strings.Split(s, "")
    for i, char := range chars {
        for k, v := range temp {
            current := make([]string, len(v[len(v) - 1]))
            copy(current, v[len(v) - 1])
            current =  append(current, char)

            temp[k][len(v)] = current
        }
        temp[i] = map[int][]string {0 : {char}}
    }

    longest := 0
    for k, v := range temp {
        for k1, v1 := range v {
            ret := isValid(v1)
            if !ret {
               delete(temp[k], k1)
               continue
            }

            if len(v1) > longest {
                fmt.Println(v1, len(v1))
                longest = len(v1)
            }
        }
    }

    return longest
}

func isValid(chars []string) bool {
    temp := []string {}

    for _, char := range chars {
        if len(temp) == 0 && char == ")" {
            return false
        }

        if char == "(" {
            temp = append(temp, char)
            continue
        }

        if char == ")" && temp[len(temp) - 1] != "(" {
            return false
        }

        if char == ")" && temp[len(temp) - 1] == "(" {
            temp = temp[:len(temp) - 1]
        }
    }

    if len(temp) > 0 {
        return false
    }

    return true
}