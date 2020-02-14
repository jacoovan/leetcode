/*
1000:合并石头的最低成本
leetcodeID : 1042
leetcode地址 : https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/
难度 : 困难

有 N 堆石头排成一排，第 i 堆中有 stones[i] 块石头。

每次移动（move）需要将连续的 K 堆石头合并为一堆，而这个移动的成本为这 K 堆石头的总数。

找出把所有石头合并成一堆的最低成本。如果不可能，返回 -1 。

 

示例 1：

输入：stones = [3,2,4,1], K = 2
输出：20
解释：
从 [3, 2, 4, 1] 开始。
合并 [3, 2]，成本为 5，剩下 [5, 4, 1]。
合并 [4, 1]，成本为 5，剩下 [5, 5]。
合并 [5, 5]，成本为 10，剩下 [10]。
总成本 20，这是可能的最小值。


示例 2：

输入：stones = [3,2,4,1], K = 3
输出：-1
解释：任何合并操作后，都会剩下 2 堆，我们无法再进行合并。所以这项任务是不可能完成的。.


示例 3：

输入：stones = [3,5,1,2,6], K = 3
输出：25
解释：
从 [3, 5, 1, 2, 6] 开始。
合并 [5, 1, 2]，成本为 8，剩下 [3, 8, 6]。
合并 [3, 8, 6]，成本为 17，剩下 [17]。
总成本 25，这是可能的最小值。


 

提示：

<ul>
    1 <= stones.length <= 30
    2 <= K <= 30
    1 <= stones[i] <= 100
</ul>

 */
package main

import(
    "fmt"
    "math"
)

func main(){
    stones := []int {3,2,4,1}
    k      := 2

    res := mergeStones(stones, k)

    fmt.Println(res)
}

func mergeStones(stones []int, K int) int {
    lenOfstone := len(stones)

    m1 := int(math.Floor(float64(lenOfstone - 1) / float64(K -1)))
    m2 := int(math.Ceil(float64(lenOfstone - 1) / float64(K -1)))

    if m1 != m2 {
        return -1
    }

    res := 0
    currentLen := lenOfstone
    for {
        if currentLen == 1 {
            break
        }

        minSum := getMinSum(stones[:currentLen], K)
        res    += minSum["sum"]

        stones[minSum["i"]] = minSum["sum"]
        if minSum["i"] + K < lenOfstone {
            i := 1
            for {
                if minSum["i"] + K + i > lenOfstone {
                    break
                }

                stones[minSum["i"] + i] = stones[minSum["i"] + K + i - 1]
                i++
            }
        }
        currentLen = currentLen - K + 1


    }

    return res
}

func getMinSum(nums []int, K int) map[string]int {
    minSum := map[string]int {
        "i" : -1,
        "sum" : -1,
    }
    lenOfnums := len(nums)

    for i, _ := range nums {
        if i + K > lenOfnums {
            break
        }

        tempSum := map[string]int {
            "i"   : i,
            "sum" : 0,
        }
        j := 0
        for {
            if j >= K {
                break
            }

            tempSum["sum"] += nums[i + j]
            j++
        }

        minSum = min(minSum, tempSum)
    }

    return minSum
}

func min(sum1 map[string]int, sum2 map[string]int) map[string]int {
    if sum1["i"] == -1 {
        return sum2
    }

    if sum1["sum"] < sum2["sum"] {
        return sum1
    }

    return sum2
}