/*
76:最小覆盖子串
leetcodeID : 76
leetcode地址 : https://leetcode-cn.com/problems/minimum-window-substring/
难度 : 困难

给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。

示例：

输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"

说明：

<ul>
	如果 S 中不存这样的子串，则返回空字符串 ""。
	如果 S 中存在这样的子串，我们保证它是唯一的答案。
</ul>

 */
package main

import(
    "fmt"
)

func main(){
    s := "ADOBECODEBANC"
    t := "ABC"
    res := minWindow(s, t)
    fmt.Println("res:", res)
}

func minWindow(s string, t string) string {
    positions := []int{}
    for i, char := range s {
        for _, needle := range t {
            if char == needle {
                positions = append(positions, i)
            }
        }
    }

    tempTMap := make(map[int]bool)
    for i, _ := range t {
        tempTMap[i] = true
    }

    windows := make([][]rune, 0)
    for _, start := range positions {
        tempTMap := make(map[int]bool)
        for i, _ := range t {
            tempTMap[i] = true
        }
        window := make([]rune, 0)
        for _, char := range s[start:] {
            window = append(window, char)
            var tempFlag = false
            for j, needle := range t {
                if char == needle {
                    tempTMap[j] = false
                }
                var tempFlag1 = false
                for _, flag := range tempTMap {
                    if flag {
                        tempFlag1 = true
                        break
                    }
                }
                tempFlag = tempFlag1
                if !tempFlag {
                    break
                }
            }
            if !tempFlag {
                windows = append(windows, window)
                break
            }
        }
    }

    if len(windows) == 0 {
        return ""
    }
    minWindowIndex := 0
    for windowIndex, window := range windows {
        if len(window) < len(windows[minWindowIndex]) {
            minWindowIndex = windowIndex
        }
    }
    minWindowSlice := windows[minWindowIndex]
    return string(minWindowSlice)
}