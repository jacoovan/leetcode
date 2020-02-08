// 请你来实现一个 atoi 函数，使其能将字符串转换成整数。

// 首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。

// 当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

// 该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。

// 注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。

// 在任何情况下，若函数不能进行有效的转换时，请返回 0。

package main

import (
    "fmt"
    "math"
)

func main() {
    s := " -4193 with words "
    s1:= atoi(s)
    fmt.Println(s1)
}

func atoi(str string) int {
    runes := []rune(str)
    fmt.Println(runes)

    noEmptyIndex := 0
    i := 0
    for {
        if runes[i] == 32 {
            noEmptyIndex++
        } else {
            break
        }
        i++
    }

    numIndex := -1
    for i,v := range runes[noEmptyIndex:]{
        if v == 45 || (v >= 48 && v <= 57) {
            numIndex = i
            break
        }
    }

    num := 0
    nums := []rune {}
    if numIndex > -1 {
        for {
            current := runes[noEmptyIndex + numIndex]
            if current == 45 || (current >= 48 && current <= 57) {
                nums = append(nums, runes[noEmptyIndex + numIndex])
            } else {
                break
            }
            numIndex++
        }
    }

    if len(nums) > 0 {
        if nums[0] == 45 {
            num = -1 * runeToInt(nums[1:])
        } else {
            num = runeToInt(nums)
        }
    }

    return num
}

func runeToInt(runes []rune) int {
    len := len(runes)
    if len == 0 {
        return 0
    }

    i, num := 0, 0
    for {
        if i >= len {
            break
        }
        current := runes[i] - 48
        num += int(math.Pow(float64(10), float64(len -i - 1)) * float64(current))
        i++
    }

    return num
}