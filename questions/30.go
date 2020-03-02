/*
30:串联所有单词的子串
leetcodeID : 30
leetcode地址 : https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/
难度 : 困难

给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

 

示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。


示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]


 */
package main

import(
    "fmt"
    "strings"
)

func main(){
    s := "barfoothefoobarman"
    words := []string {"foo", "bar"}

    res := findSubstring(s, words)

    fmt.Println(res)
}

type WordRes struct {
    word string
    warr []string
    used bool
    finished bool
    index int
    start int
}

func findSubstring(s string, words []string) []int {
    wMap := initWords(words)
    res := []int {}

    tempArr := strings.Split(s, "")

    i := 0
    tempPosition := 0
    for {
        position := tempPosition
        sArr := tempArr

        i++
        if i == 100 {
            return res
        }
        sArr, ok, position := getNext(sArr, wMap, position)
        _ = sArr

        if ok {
            if isFinished(wMap) {
                res = append(res, getStart(wMap))
            }
        } else {
            wMap = initWords(words)
        }


        if len(sArr) == 0 {
            return res
        }
        tempArr = sArr
        tempPosition = position

    }
}

func getStart(wMap map[string]WordRes) int {
    res := -1

    for _, wres := range wMap {
        if res == -1 {
            res = wres.start
        } else if wres.start < res {
            res = wres.start
        }
    }

    return res
}

func isFinished(wMap map[string]WordRes) bool {
    for _, wres := range wMap {
        if wres.finished == false {
            return false
        }
    }

    return true
}

func getMatch(sArr []string, dArr []string) bool {
    for index, v1 := range dArr {
        if v1 != sArr[index] {
            return false
        }
    }
    return true
}

func getNext(sArr []string, wMap map[string]WordRes, position int) ([]string, bool, int) {
    matched := false

    for word, wres := range wMap {
        if wres.used == false {
            if sArr[0] == wres.warr[0] {
                matched = getMatch(sArr, wres.warr)
                wres.finished = true
                wres.used = true
                wres.start = position
                wMap[word] = wres
                index := len(wres.word)
                sArr = sArr[index : ]
                position = position + index
                break
            }
        }
    }

    if matched == false && len(sArr) > 0 {
        sArr = sArr[1:]
        position++
    }

    return sArr, matched, position
}

func initWords(words []string) map[string]WordRes {
    var res = map[string]WordRes {}

    for _, word := range words {
        res[word] = WordRes {
            word,
            strings.Split(word, ""),
            false,
            false,
            0,
            0,
        }
    }

    return res
}