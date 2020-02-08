// 给定一个排序数组，你需要使用·原地算法·删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度
// 不需要使用额外的数组空间，你必须使用原地算法修改输入数组并在使用O(1)额外空间的条件下完成

package main

import "fmt"

func main() {
    arr := []int {1,1,1,1,2}
    index, arr := array_unique(arr)
    fmt.Println(index, arr)
}

func array_unique(arr []int) (int, []int) {
    index := 0
    len   := len(arr)

    if len == 0 {
        return 0, []int {}
    }

    i := 0
    for {
        if i >= len - 1 {
            break
        }

        if arr[i] == arr[i+1] {
            i++
            continue
        }

        if i != index {
            arr[i + 1], arr[index + 1] = arr[index + 1], arr[i + 1]
            index++
        }
        i++
    }

    return index + 1, arr
}