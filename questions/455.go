/*
455:分发饼干
leetcodeID : 455
leetcode地址 : https://leetcode-cn.com/problems/assign-cookies/
难度 : 简单

假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。对每个孩子 i ，都有一个胃口值 g<sub>i ，</sub>这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j ，都有一个尺寸 s<sub>j </sub>。如果 s<sub>j</sub> >= g<sub>i </sub>，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

注意：

你可以假设胃口值为正。<br />
一个小朋友最多只能拥有一块饼干。

示例 1:


输入: [1,2,3], [1,1]

输出: 1

解释: 
你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
所以你应该输出1。


示例 2:


输入: [1,2], [1,2,3]

输出: 2

解释: 
你有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
你拥有的饼干数量和尺寸都足以让所有孩子满足。
所以你应该输出2.


 */
package main

import(
    "fmt"
)

func main(){
    childs  := []int {1,2}
    cookies := []int {1,2}

    res := findContentChildren(childs, cookies)

    fmt.Println(res)
}

func findContentChildren(g []int, s []int) int {
    ate := []int {}

    for _, gi := range g {
        for j, sj := range s {
            if gi == sj && !inArray(j, ate) {
                ate = append(ate, j)
            }
        }
    }

    return len(ate)
}

func inArray(n int, nums []int) bool {
    for _, n1 := range nums {
        if n == n1  {
            return true
        }
    }
    return false
}