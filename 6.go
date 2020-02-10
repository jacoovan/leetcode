/*
6:Z 字形变换
leetcodeID : 6
leetcode地址 : https://leetcode-cn.com/problems/zigzag-conversion/
难度 : 中等

将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N


之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);

示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"


示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

 */
package main

import(
    "fmt"
)

func main(){
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