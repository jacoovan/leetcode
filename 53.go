// 给定一个整数数组·nums·,找到一个具有最大和的联系子数组（子数组至少包含一个元素）,返回其最大和

package main

import "fmt"

func main() {
    nums := []int {1,4,7,8,2}
    max  := get_maxsum(nums)
    fmt.Println(max)
}

func get_maxsum(arr []int) int {
    maxsum, sum := -1 << 31, -1 << 31
    for _, num := range arr {
        sum := max(sum + num, num)
        maxsum = max(maxsum, sum)
    }
    maxsum = maxsum

    return maxsum
}

func max(num1 int, num2 int) int {
    if num1 < num2 {
        return num2
    }
    return num1
}