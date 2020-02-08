// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

package main

import (
    "fmt"
)

func main() {
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