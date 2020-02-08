// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
package main

import (
    "fmt"
)

func main() {
    nums := []int {1,2,3}

    ret  := calc(nums)
    
    fmt.Println(ret)
}

func calc(nums []int) []int {
    len := len(nums)
    if len == 0 {
        return []int {}
    }

    add := (bool) (nums[0] == 9)

    var i int
    for i = len - 1; i >= 0; i-- {
        if nums[i] < 9 {
            nums[i] = nums[i] + 1
            break
        } else {
            nums[i] = 0
        }
    }

    if add && i == -1 {
        ret := make([]int, len + 1)
        ret[0] = 1
        for i, v := range nums {
            ret[i + 1] = v
        }
        nums = ret
    }

    return nums
}