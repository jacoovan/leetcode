/*
79:单词搜索
leetcodeID : 79
leetcode地址 : https://leetcode-cn.com/problems/word-search/
难度 : 中等

给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中&ldquo;相邻&rdquo;单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:

board =
[
  [&#39;A&#39;,&#39;B&#39;,&#39;C&#39;,&#39;E&#39;],
  [&#39;S&#39;,&#39;F&#39;,&#39;C&#39;,&#39;S&#39;],
  [&#39;A&#39;,&#39;D&#39;,&#39;E&#39;,&#39;E&#39;]
]

给定 word = "ABCCED", 返回 true.
给定 word = "SEE", 返回 true.
给定 word = "ABCB", 返回 false.

*/
package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCB"
	res := exist(board, word)
	fmt.Println("res:", res)
}

type Position struct {
	X int
	Y int
}

func exist(board [][]byte, word string) bool {
	stack := make([]Position, 0)
	existMaps := make([]map[int]map[int]bool, 0)
	for i, row := range board {
		for j, item := range row {
			if item == word[0] {
				stack = append(stack, Position{X: j, Y: i})
				existMap := make(map[int]map[int]bool)
				for x := 0; x < len(row); x++ {
					existMap[x] = make(map[int]bool)
				}
				existMap[j][i] = true
				existMaps = append(existMaps, existMap)
			}
		}
	}
	if len(stack) == 0 {
		return false
	}

	for i, pos := range stack {
		if seek(board, existMaps[i], pos, word, 1) {
			return true
		}
	}
	return false
}

func seek(board [][]byte, existMap map[int]map[int]bool, position Position, word string, index int) bool {
	if index == len(word) {
		return true
	}
	stack := make([]Position, 0)
	existMaps := make([]map[int]map[int]bool, 0)
	positions := make([]Position, 0)
	if position.Y-1 >= 0 && !existMap[position.X][position.Y-1] {
		e := newExistMap(existMap, len(board[0]))
		e[position.X][position.Y-1] = true
		existMaps = append(existMaps, e)
		positions = append(positions, Position{position.X, position.Y - 1})
	}
	if position.Y+1 <= len(board)-1 && !existMap[position.X][position.Y+1] {
		e := newExistMap(existMap, len(board[0]))
		e[position.X][position.Y+1] = true
		existMaps = append(existMaps, e)
		positions = append(positions, Position{position.X, position.Y + 1})
	}
	if position.X-1 >= 0 && !existMap[position.X-1][position.Y] {
		e := newExistMap(existMap, len(board[0]))
		e[position.X-1][position.Y] = true
		existMaps = append(existMaps, e)
		positions = append(positions, Position{position.X - 1, position.Y})
	}
	if position.X+1 <= len(board[0])-1 && !existMap[position.X+1][position.Y] {
		e := newExistMap(existMap, len(board[0]))
		e[position.X+1][position.Y] = true
		existMaps = append(existMaps, e)
		positions = append(positions, Position{position.X + 1, position.Y})
	}
	for i, pos := range positions {
		if board[pos.Y][pos.X] != word[index] {
			continue
		}
		stack = append(stack, Position{pos.X, pos.Y})
		existMaps[i][pos.X][pos.Y] = true
		length := 0
		for _, row := range existMaps[i] {
			length += len(row)
		}

		if length == len(word) {
			return true
		}
		if seek(board, existMaps[i], pos, word, index+1) {
			return true
		}
	}

	return false
}

func newExistMap(existMap map[int]map[int]bool, width int) map[int]map[int]bool {
	e := make(map[int]map[int]bool)
	for x := 0; x < width; x++ {
		e[x] = make(map[int]bool)
	}
	for x, row := range existMap {
		for y, _ := range row {
			e[x][y] = true
		}
	}
	return e
}
