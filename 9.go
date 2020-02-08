// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

package main

import (
    "fmt"
)

func main() {
    num := -121
    res := isPalindrome(num)
    fmt.Println(res)
}

func isPalindrome(num int) bool {
    if num < 0 {
        return false
    }

    len := 0
    tmp, nums := 0, []int {}
    for {
        if float64(num) / 10 > 0 {
            tmp  = num - int(num / 10) * 10
            nums = append(nums, tmp) 
            num = int(num / 10)
            len++
        } else {
            break
        }
    }

    i, res := 0, true
    for {
        if i > len - 1 {
            break
        }

        if nums[i] != nums[len - 1 - i] {
            res = false
            break
        }

        i++
    }
    return res
}