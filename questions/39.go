/*
39:组合总和
leetcodeID : 39
leetcode地址 : https://leetcode-cn.com/problems/combination-sum/
难度 : 中等

给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

<ul>
	所有数字（包括 target）都是正整数。
	解集不能包含重复的组合。 
</ul>

示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]


示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

 */
package main

import (
    "fmt"
    "math"
    "sort"
    "strconv"
)

func main(){
    nums := []int {1,3,5}
    target := 9

    res := combinationSum(nums, target)
    fmt.Println(res)
}

func combinationSum(candidates []int, target int) [][]int {
    answer := getAnswer(candidates, target)

    uniqueAnswer := answerUnique(answer)

    return uniqueAnswer
}

func answerUnique(answer [][]int) [][]int {
    unique := map[string][]int {}

    for _, v := range answer {
        sort.Ints(v)
        key := ""
        for _, num := range v {
            key += strconv.Itoa(num)
        }

        _, ok := unique[key]
        if !ok {
            unique[key] = v
        }
    }

    res := [][]int {}
    for _, v := range unique {
        res = append(res, v)
    }

    return res
}

func getAnswer(candidates []int, target int) [][]int {
    res := [][]int {}

    for i, num := range candidates {
        size := int(math.Floor(float64(target) / float64(num)))
        for j := 1; j <= size; j++ {
            tempNums := make([]int, i)
            copy(tempNums, candidates[:i])
            if i < size {
                tempNums = append(tempNums, candidates[i+1:]...)
            }

            current := []int {}
            for k := 0; k < j; k++ {
                current = append(current, num)
            }


            if target - num * j == 0 {
                res = append(res, current)
            }
            if target - num * j > 0 {
                res1 := combinationSum(tempNums, target - num * j)

                for k, row := range res1 {
                    sum := 0
                    for _, v := range row {
                        sum += v
                    }

                    for _, v := range current {
                        sum += v
                    }

                    if sum == target {
                        res1[k] = append(current, res1[k]...)
                    }
                }

                res = append(res, res1...)
            }


        }
    }

    return res
}