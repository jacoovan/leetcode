// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

package main

import (
    "fmt"
    "strconv"
    "math"
)

func main () {
    nums1 := []int {1,2}
    nums2 := []int {3,4}

    middleNum := getMiddleNum(nums1, nums2)

    fmt.Println(middleNum)
}

func getMiddleNum(nums1 []int, nums2 []int) float64 {
    len1, len2 := len(nums1), len(nums2)
    len := len(nums1) + len(nums2)
    sinleNum := false
    if len % 2 > 0 {
        sinleNum = true
    }
    middleIndex, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(len) / 2), 64)
    middleIndex = math.Ceil(middleIndex)

    i, j, ret := 0, 0, 0.0

    for {
        exist1 := false
        if i < len1 -1 && nums1[i] <= nums2[j] {
            i++
            exist1 = true
        } else if j < len2 -1 {
            j++
        }

        if sinleNum {
            if int64(i + j) == int64(middleIndex) {
                if exist1 {
                    ret = float64(nums1[i - 1])
                    break
                } else {
                    ret = float64(nums2[j - 1])
                    break
                }
            }
        } else {
            if (i + j) == len / 2 {
                if exist1 {
                    i++
                } else {
                    j++
                }
                ret = (float64(nums1[i - 1]) + float64(nums2[j - 1])) / 2
                break
            }
        }
    }
    
    return ret
}