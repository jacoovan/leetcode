// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在返回它将会被按顺序插入的位置

package main

import "fmt"

func main() {
    arr   := []int {1,3,4,5,6,7,9}
    index := search_item(arr, 2)
    fmt.Println(index)
}

func search_item(arr []int, target_num int) int {
    i   := 0
    len := len(arr)
    for {
        if arr[i] >= target_num {
            return i
        }
        i++
        if i > len {
            break
        }
    }
    return -1
}