// 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

package main

import (
    "fmt"
)

func main() {
    s := "LEETCODEISHIRING"
    s1:= convert(s, 3)
    fmt.Println(s1)
}

func convert(s string, numRows int) string {
    words := []rune(s)

    zmap := make([][]rune, numRows)
    numCols := len(words)
    for i, _ := range zmap {
        zmap[i] = make([]rune, numCols)
    }

    lastPosition := map[string]int {
        "i"   : -1,
        "row" : -1,
        "col" : -1,
    }
    completeZmap(words, zmap, lastPosition)

    var res []rune
    for _, v := range zmap {
        for _, v1 := range v {
            if v1 > 0 {
                res = append(res, v1)
            }
        }
    }

    return string(res)
}

func completeZmap(words []rune, zmap [][]rune, lastPosition map[string]int) bool {
    numRows := len(zmap)

    lastI := lastPosition["i"]
    lastR := lastPosition["row"]
    lastC := lastPosition["col"]

    if lastI == len(words) - 1 {
        return true
    }

    row, col := 0, 0
    if lastR == -1 || lastC == -1 {
        row, col = 0, 0
    } else if lastR == numRows - 1 {
        row = numRows - 2
        col = lastC + 1
    } else if lastR == 0 {
        row = 1
        col = lastC
    } else if lastC % (numRows - 1) > 0 {
        row = lastR - 1 
        col = lastC + 1
    } else if lastC % (numRows - 1) == 0 {
        row = lastR + 1
        col = lastC
    }

    zmap[row][col] = words[lastI + 1]

    lastPosition["i"]   = lastI + 1
    lastPosition["row"] = row
    lastPosition["col"] = col

    completeZmap(words, zmap, lastPosition)

    return true
}