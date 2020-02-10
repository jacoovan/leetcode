/*
12:整数转罗马数字
leetcodeID : 12
leetcode地址 : https://leetcode-cn.com/problems/integer-to-roman/
难度 : 中等

罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

<ul>
	I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
	X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
	C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
</ul>

给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。

示例 1:

输入: 3
输出: "III"

示例 2:

输入: 4
输出: "IV"

示例 3:

输入: 9
输出: "IX"

示例 4:

输入: 58
输出: "LVIII"
解释: L = 50, V = 5, III = 3.


示例 5:

输入: 1994
输出: "MCMXCIV"
解释: M = 1000, CM = 900, XC = 90, IV = 4.

 */
package main

import(
    "fmt"
    "math"
)

func main(){
    num := 1994
    res := intToRoman(num)
    fmt.Println(res)
}

func intToRoman(num int) string {
    nums := []int {}

    for {
        if float64(num) / 10 > 0 {
            num1 := int(math.Floor(float64(num / 10)))
            temp := num - num1 * 10
            num  =  num1
            nums =  append(nums, temp)
        } else {
            break
        }
    }

    var zero string
    var one  string
    var five string
    var ten  string
    var res  []string
    for i, v := range nums {
        switch i {
            case 0:
                one  = "I"
                five = "V"
                ten  = "X"
                break
            case 1:
                zero = "X"
                one  = "X"
                five = "L"
                ten  = "C"
                break
            case 2:
                zero = "C"
                one  = "C"
                five = "D"
                ten  = "M"
                break
            case 3:
                one  = "M"
                break
        }

        temp := ""
        if v == 0 {
            temp = zero
        } else if v < 4 {
            for {
                if v <= 0 {
                    break
                }
                temp += one
                v--
            }
        } else if v == 4 {
            temp = one + five
        } else if v == 5 {
            temp = five
        } else if v > 5 && v < 9 {
            temp += five
            for {
                if v <= 5 {
                    break
                }
                temp += one
                v--
            }
        } else if v == 9 {
            temp = one + ten
        }
        res = append(res, temp)
    }

    ret := ""
    i   := len(res) - 1
    for {
        if i < 0 {
            break
        }
        ret += res[i]
        i--
    }

    return ret
}