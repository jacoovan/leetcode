/*
17:电话号码的字母组合
leetcodeID : 17
leetcode地址 : https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
难度 : 中等

给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

<img src="https://assets.leetcode-cn.com/aliyun-lc-upload/original_images/17_telephone_keypad.png" style="width: 200px;">

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].


说明:<br>
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

 */
package main

import (
    "fmt"
)
func main(){
    digits := "**"
    res := letterCombinations(digits)
    fmt.Println(res, len(res))
}

func letterCombinations(digits string) []string {
    res := []string{}
    numLetterMap := map[rune][]string {
        '2': {"a", "b", "c"},
        '3': {"d", "e", "f"},
        '4': {"g", "h", "i"},
        '5': {"j", "k", "l"},
        '6': {"m", "n", "o"},
        '7': {"p", "q", "r", "s"},
        '8': {"t", "u", "v"},
        '9': {"w", "x", "y", "z"},
    }
    for _, s := range digits {
        letters, _ := numLetterMap[s]
        if len(res) == 0 {
            for _, letter := range letters {
                res = append(res, letter)
            }
            continue
        }

        temp := []string{}
        for _, word := range res {
            for _, letter := range letters {
                temp = append(temp, word+letter)
            }
        }
        res = temp
    }
    return res
}

