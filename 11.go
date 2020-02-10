/*
11:盛最多水的容器
leetcodeID : 11
leetcode地址 : https://leetcode-cn.com/problems/container-with-most-water/
难度 : 中等

给定 n 个非负整数 a<sub>1</sub>，a<sub>2，</sub>...，a<sub>n，</sub>每个数代表坐标中的一个点 (i, a<sub>i</sub>) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, a<sub>i</sub>) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

<img alt="" src="https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg" style="height: 287px; width: 600px;">

<small>图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。</small>

 

示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49

 */
package main

import(
    "fmt"
)

func main(){
    nums := []int {1,8,6,2,5,4,8,3,7}

    area := maxArea(nums)

    fmt.Println(area)
}

func maxArea(height []int) int {
    len := len(height)

    if len < 2 {
        return 0
    }

    var area int
    var tmpArea int
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

            tmpArea = int(j - i) * min(height[i], height[j])
            if tmpArea > area {
                area = tmpArea
            }

            j++
        }

        i++
    }

    return area
}

func min(i int, j int) int {
    if i > j {
        return j
    }
    return i
}