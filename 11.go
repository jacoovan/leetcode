// 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

// 说明：你不能倾斜容器，且 n 的值至少为 2。

package main

import (
    "fmt"
)

func main() {
    nums := []float64 {1,8,6,2,5,4,8,3,7}

    area := maxArea(nums)

    fmt.Println(area)
}

func maxArea(nums []float64) float64 {
    len := len(nums)

    if len < 2 {
        return 0
    }

    var area float64
    var tmpArea float64
    area, tmpArea = 0, 0
    i, j := 0, 1
    for {
        if i > len - 2 {
            break
        }

        j = i + 1

        for {
            if j > len - 1 {
                break
            }

            tmpArea = float64(j - i) * min(nums[i], nums[j])
            if tmpArea > area {
                area = tmpArea
            }

            fmt.Println(i, j, nums[i], nums[j], tmpArea)

            j++
        }

        i++
    }

    return area
}

func min(i float64, j float64) float64 {
    if i > j {
        return j
    }
    return i
}