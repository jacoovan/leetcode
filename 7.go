// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

package main

import (
    "fmt"
    "math"
)

func main() {
    n := -123
    n1:= reverse(n)
    fmt.Println(n1)
}

func reverse(x int) int {
    if x == -2 ^ 31 {
        return 0
    }

    nums := []int {}
    nums = r1(x, nums)
    res := 0
    if len(nums) > 0 {
        for i := 0; i <= len(nums) - 1; i++ {
            res += nums[i] * int(math.Pow(10, float64(len(nums) - 1 - i)))
        }
    }
    return res
}

func r1(x int, nums []int) []int{
    if math.Abs(float64(x)) < 1 {
        return nums
    }

    nums = append(nums, x % 10)
    x    = int((x - x % 10) / 10)

    nums = r1(x, nums)

    return nums
}