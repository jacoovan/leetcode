/*
40:组合总和 II
leetcodeID : 40
leetcode地址 : https://leetcode-cn.com/problems/combination-sum-ii/
难度 : 中等

给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

<ul>
	所有数字（包括目标数）都是正整数。
	解集不能包含重复的组合。 
</ul>

示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]


示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]

 */
package main

import(
    "fmt"
    "sort"
    "strconv"
)

func main(){
    nums := []int {10,1,2,7,6,1,5}
    target := 8

    res := combinationSum2(nums, target)
    fmt.Println(res)
}

func combinationSum2(candidates []int, target int) [][]int {
    answer := getAnswer2(candidates, target)

    uniqueAnswer := answerUnique2(answer)

    return uniqueAnswer
}

func answerUnique2(answer [][]int) [][]int {
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

func getAnswer2(candidates []int, target int) [][]int {
    res := [][]int {}

    for i, num := range candidates {
        size := 1
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
                res1 := combinationSum2(tempNums, target - num * j)

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