/*
396:旋转函数
leetcodeID : 396
leetcode地址 : https://leetcode-cn.com/problems/rotate-function/
难度 : 中等

给定一个长度为 n 的整数数组 A 。

假设 B<sub>k</sub> 是数组 A 顺时针旋转 k 个位置后的数组，我们定义 A 的&ldquo;旋转函数&rdquo; F 为：

F(k) = 0 * B<sub>k</sub>[0] + 1 * B<sub>k</sub>[1] + ... + (n-1) * B<sub>k</sub>[n-1]。

计算F(0), F(1), ..., F(n-1)中的最大值。

注意:<br />
可以认为 n 的值小于 10<sup>5</sup>。

示例:


A = [4, 3, 2, 6]

F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26

所以 F(0), F(1), F(2), F(3) 中的最大值是 F(3) = 26 。


 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func maxRotateFunction(A []int) int {
    
}