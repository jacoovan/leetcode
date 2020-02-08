// 给定一个整数数组和一个目标值，找出数组中和为目标值的两个数
// 你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用

package main

import "fmt"

func main () {
    target := arr_sum([]int{0,1,2}, 3)
    fmt.Println(target)
}

func arr_sum(arr []int, target_num int) map[int]int {

    var arrmap map[int]int
    arrmap =  make(map[int]int)
    length := len(arr)
    i      := 0
    for {
        if i >= length {
            break
        }
        if arr[i] != target_num {
            arrmap[i] = arr[i]
        }

        for j, num := range arrmap {
            fmt.Println(j)
            if i != j && num + arr[i] == target_num {
                return map[int]int {
                    j : num,
                    i : arr[i],
                }
            }
        }
        i++
    }

    return nil
}