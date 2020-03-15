/*
43:字符串相乘
leetcodeID : 43
leetcode地址 : https://leetcode-cn.com/problems/multiply-strings/
难度 : 中等

给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"

示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"

说明：


	num1 和 num2 的长度小于110。
	num1 和 num2 只包含数字 0-9。
	num1 和 num2 均不以零开头，除非是数字 0 本身。
	不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。


 */
package main

import (
    "fmt"
    "math"
    "strconv"
    "strings"
)

func main(){
    num1 := "123"
    num2 := "456"
    res := multiply(num1, num2)

    fmt.Println(res)
}

func multiply(num1 string, num2 string) string {
    nums1 := strings.Split(num1, "")
    nums2 := strings.Split(num2, "")

    tempResult := map[int]int {}
    pos := 0
    for i, first := range nums1 {
        for j, second := range nums2 {
            n1, _ := strconv.Atoi(first)
            n2, _ := strconv.Atoi(second)


            r := n1 * n2
            r0 := r - int(math.Floor(float64(r) / 10.0)) * 10
            r1 := int(math.Floor(float64(r) / 10.0))

            pos = i + j

            v, ok := tempResult[pos]
            if !ok {
                tempResult[pos] = r0
            } else if v + r0 < 10 {
                tempResult[pos] += r0
            } else {
                tempResult[pos] += r0
                tempResult[pos] -= 10
                r1 += 1
            }

            if r1 > 0 {
                v1, _ := tempResult[pos - 1]
                if v1 + r1 < 10 {
                    tempResult[pos - 1] += r1
                } else {
                    tempResult[pos - 1] += r1
                    tempResult[pos - 1] -= 10

                    for {
                        pos  = pos - 1
                        v2, ok := tempResult[pos - 1]
                        if !ok {
                            v2 = 1
                            tempResult[pos - 1] = v2
                            break
                        } else if v2 + 1 < 10 {
                            v2 += 1
                            tempResult[pos - 1] = v2
                            break
                        } else {
                            v2 = 0
                            tempResult[pos - 1] = 0
                        }
                    }
                }
            }

        }
    }

    min := 0
    for i, _ := range tempResult {
       if i < min {
           min = i
       }
    }
    res := ""
    for i := min; i < min + len(tempResult); i++ {
        res += strconv.Itoa(tempResult[i])
    }

    return res
}