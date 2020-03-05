/*
36:有效的数独
leetcodeID : 36
leetcode地址 : https://leetcode-cn.com/problems/valid-sudoku/
难度 : 中等

判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。


	数字 1-9 在每一行只能出现一次。
	数字 1-9 在每一列只能出现一次。
	数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。


<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/sudoku-by-l2g-20050714.svg/250px-sudoku-by-l2g-20050714.svg.png" style="height: 250px; width: 250px;">

<small>上图是一个部分填充的有效的数独。</small>

数独部分空格内已填入了数字，空白格用 &#39;.&#39; 表示。

示例 1:

输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true


示例 2:

输入:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: false
解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
     但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。

说明:

<ul>
	一个有效的数独（部分已被填充）不一定是可解的。
	只需要根据以上规则，验证已经填入的数字是否有效即可。
	给定数独序列只包含数字 1-9 和字符 &#39;.&#39; 。
	给定数独永远是 9x9 形式的。
</ul>

 */
package main

import(
    "fmt"
)

func main(){
    board := [][]byte {
        {'5','3','.','.','7','.','.','.','.'},
        {'6','.','.','1','9','5','.','.','.'},
        {'.','9','8','.','.','.','.','6','.'},
        {'8','.','.','.','6','.','.','.','3'},
        {'4','.','.','8','.','3','.','.','1'},
        {'7','.','.','.','2','.','.','.','6'},
        {'.','6','.','.','.','.','2','8','.'},
        {'.','.','.','4','1','9','.','.','5'},
        {'.','.','.','.','8','.','.','7','9'},
    }

    res := isValidSudoku(board)

    fmt.Println(res)
}

func isValidSudoku(board [][]byte) bool {
    for _, row := range board {
        res := rowOrcolUnique(row)
        if !res {
            return false
        }
    }

    for i := 0; i < len(board[0]); i++ {
        col := []byte {}
        for _, row := range board {
            col = append(col, row[i])
        }

        res := rowOrcolUnique(col)

        if !res {
            return false
        }
    }

    for i := 0; i < len(board) - 3; i=i+3 {
        for j := 0; j < len(board[0]) - 3; j=j+3 {
            boardChunk := make([][]byte, 3)
            copy(boardChunk, board[i:i+3])
            for i, v := range boardChunk {
                boardChunk[i] = v[j:j+3]
            }

            res := boardUnique(boardChunk)
            if !res {
                return false
            }
        }
    }
    return true
}

func rowOrcolUnique(row []byte) bool {
    temp := map[byte]bool {}

    for _, v := range row {
        _, ok := temp[v]
        if !ok {
            temp[v] = true
        } else if v != '.' {
            return false
        }
    }
    return true
}

func boardUnique(board [][]byte) bool {
    temp := map[byte]bool {}

    for _, row := range board {
        for _, v := range row {
            _, ok := temp[v]
            if !ok {
                temp[v] = true
            } else if v != '.' {
                return false
            }
        }
    }
    return true
}