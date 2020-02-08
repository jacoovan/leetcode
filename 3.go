// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

package main

import (
    "fmt"
)

func main () {
    s := "pwwkew"
    l := lengthOfLongestSubString(s)
    fmt.Println(l)
}

func lengthOfLongestSubString(s string) int {
    runes := []rune(s)
    var current map[string]string
    current = make(map[string]string)
    // current = map[string]string {}

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