/*
59:螺旋矩阵 II
leetcodeID : 59
leetcode地址 : https://leetcode-cn.com/problems/spiral-matrix-ii/
难度 : 中等

给定一个正整数 n，生成一个包含 1 到 n<sup>2</sup> 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]

 */
package main

import(
    "fmt"
    "math"
)

func main(){
    n := 16

    mat := generateMatrix(n)

    printMat(mat)
}

func generateMatrix(n int) [][]int {
    sqrt := int(math.Ceil(math.Sqrt(float64(n))))
    mat := [][]int {}
    for i := 1; i <= sqrt; i++ {
        line := []int {}
        for j := 1; j <= sqrt; j++ {
            line = append(line, 0)
        }
        mat = append(mat, line)
    }

    times := 0
    num := 1
    for k := 1; k <= sqrt; k++ {
        fmt.Println("times", times)
        row, col := times, times
        end := sqrt - times - 1
        for ; col <= end; col++ {
            mat[row][col] = num
            num++
            if num > n {return mat}
        }
        printMat(mat)

       if sqrt - times == 2 {
           break
       }
        row, col = times + 1, sqrt - times - 1
        end = sqrt - times - 1
        for ; row <= end; row++ {
            mat[row][col] = num
            num++
            if num > n {return mat}
        }
        printMat(mat)

        row, col = sqrt - times - 1, sqrt - times - 2
        for ; col >= times; col-- {
            mat[row][col] = num
            num++
            if num > n {return mat}
        }
        printMat(mat)

        row, col = sqrt - times - 2, times
        for ; row >= times + 1; row-- {
            mat[row][col] = num
            num++
            if num > n {return mat}
        }
        printMat(mat)
        times++
    }
    return mat
}

func printMat(mat [][]int) {
    fmt.Println("")
    for i := 0; i <= len(mat) - 1; i++ {
        fmt.Println(mat[i])
    }
}