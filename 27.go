// 给定一个数组nums和一个值val,你需要使用·原地算法·移除所有数值等于val的元素,返回移除后数组的新长度
// 不需要使用额外的数组空间，你必须使用·原地算法修改输入数组·并在O(1)额外空间的条件下完成
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素

package main

import "fmt"

func main() {
    arr := []int{1,1,1,3}
    index, arr := remove_duplicate(arr, 1)
    fmt.Println(index, arr)
}

func remove_duplicate(arr []int, target_num int) (int, []int) {
    i := 0
    j := len(arr) - 1
    for {
        if i >= j {
            break
        }

        if arr[i] == target_num {
            for {
                if arr[j] == target_num {
                    j--
                    continue
                }
                break
            }

            arr[i], arr[j] = arr[j], arr[i]
            i++
        } else if arr[i] != target_num && i < j {
            i++
        } else {
            break
        }
    }

    return i, arr
}