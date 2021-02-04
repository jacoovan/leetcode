/*
474:一和零
leetcodeID : 474
leetcode地址 : https://leetcode-cn.com/problems/ones-and-zeroes/
难度 : 中等

在计算机界中，我们总是追求用有限的资源获取最大的收益。

现在，假设你分别支配着 m 个 0 和 n 个 1。另外，还有一个仅包含 0 和 1 字符串的数组。

你的任务是使用给定的 m 个 0 和 n 个 1 ，找到能拼出存在于数组中的字符串的最大数量。每个 0 和 1 至多被使用一次。

注意:


	给定 0 和 1 的数量都不会超过 100。
	给定字符串数组的长度不会超过 600。


示例 1:


输入: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
输出: 4

解释: 总共 4 个字符串可以通过 5 个 0 和 3 个 1 拼出，即 "10","0001","1","0" 。


示例 2:


输入: Array = {"10", "0", "1"}, m = 1, n = 1
输出: 2

解释: 你可以拼出 "10"，但之后就没有剩余数字了。更好的选择是拼出 "0" 和 "1" 。


*/
package main

import "fmt"

func main() {
	strs := []string{"10", "0", "1"}
	m := 1
	n := 1

	// strs := []string{"10", "0001", "111001", "1", "0"}
	// m := 5
	// n := 3

	res := findMaxForm(strs, m, n)

	fmt.Println("res:", res)
}

func findMaxForm(strs []string, m int, n int) int {
	tempList := make([]*SortObject, 0)
	for i := range strs {
		tempList = append(tempList, NewSortObject(&strs[i]))
	}
	sortFast(&tempList)

	count := 0
	for _, object := range tempList {
		isBreak := false
		tempM := m
		tempN := n
		for _, c := range *object.origin {
			if c == '0' {
				tempM = tempM - 1
				if tempM < 0 {
					isBreak = true
					break
				}
			}

			if c == '1' {
				tempN = tempN - 1
				if tempN < 0 {
					isBreak = true
					break
				}
			}
		}
		if !isBreak {
			count = count + 1
			m = tempM
			n = tempN
		}
	}
	return count
}

func sortFast(objectsPtr *[]*SortObject) {
	if len(*objectsPtr) == 0 {
		return
	}
	originObjects := *objectsPtr
	smaller := make([]*SortObject, 0)
	bigger := make([]*SortObject, 0)
	flag := originObjects[0]
	for _, object := range originObjects[1:] {
		if object.Len() >= flag.Len() {
			bigger = append(bigger, object)
		} else {
			smaller = append(smaller, object)
		}
	}

	if len(bigger) > 1 {
		sortFast(&bigger)
	}
	if len(smaller) > 1 {
		sortFast(&smaller)
	}

	index := 0
	for _, small := range smaller {
		originObjects[index] = small
		index = index + 1
	}
	originObjects[index] = flag
	index = index + 1
	for _, big := range bigger {
		originObjects[index] = big
		index = index + 1
	}
	return
}

type SortObject struct {
	length int
	origin *string
}

func NewSortObject(strPtr *string) *SortObject {
	return &SortObject{
		length: len(*strPtr),
		origin: strPtr,
	}
}

func (s *SortObject) Len() int {
	return s.length
}
