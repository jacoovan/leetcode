/*
42:接雨水
leetcodeID : 42
leetcode地址 : https://leetcode-cn.com/problems/trapping-rain-water/
难度 : 困难

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

<img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png" style="height: 161px; width: 412px;">

<small>上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。</small>

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

 */
package main

import(
    "fmt"
)

func main(){
    height := []int {0,1,0,2,1,0,1,3,2,1,2,1}

    res := trap(height)

    fmt.Println(res)
}

func trap(height []int) int {
    barrels := getBarrels(height)

    volume := 0
    for _, barrel := range barrels {
        volume += getVolume(barrel, height)
    }

    return volume
}

func getBarrels(height []int) [][]int {
    if len(height) == 0 {
        return [][]int {}
    }

    first := getFirstMax(height)
    second := 0
    if first + 1 < len(height) {
        second = getSecondMax(height[first:])
        if second == 0 {
            return [][]int {}
        }
        second = first + second
    } else {
        return [][]int {}
    }

    barrels := [][]int {
        {first, second},
    }
    if second + 1 < len(height) {

        nextBarrels := getBarrels(height[second:])

        if len(nextBarrels) > 0 {
            for _, barrel := range nextBarrels {
                barrel[0] += second
                barrel[1] += second
                barrels = append(barrels, barrel)
            }
        }

        return barrels
    } else {
        return [][]int {}
    }
}

func getFirstMax(height []int) int {
    maxIndex := 0

    for i := 1; i < len(height); i++ {
        if height[i] < height[maxIndex] {
            break
        } else {
            maxIndex = i
        }
    }

    return maxIndex
}

func getSecondMax(height []int) int {
    maxIndex := 0

    sub := true
    for i := 1; i < len(height); i++ {
        if sub && height[i] <= height[i - 1] {
            continue
        }

        if sub && height[i] > height[i - 1] {
            sub = false
            continue
        }

        if !sub && height[i - 1] >= height[i] {
            maxIndex = i - 1
            break
        }

        // 未完待续...
    }

    return maxIndex
}

func getVolume(barrel []int, height []int) int {
    start := barrel[0]
    end   := barrel[1]

    maxHeight := 0
    if height[start] > height[end] {
        maxHeight = height[end]
    } else {
        maxHeight = height[start]
    }

    volumn := 0
    for i := start; i <= end; i++ {
        if height[i] > maxHeight {
            continue
        }
        volumn += maxHeight - height[i]
    }

    return volumn
}