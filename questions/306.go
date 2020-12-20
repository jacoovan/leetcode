/*
306:累加数
leetcodeID : 306
leetcode地址 : https://leetcode-cn.com/problems/additive-number/
难度 : 中等

累加数是一个字符串，组成它的数字可以形成累加序列。

一个有效的累加序列必须至少包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。

给定一个只包含数字 &#39;0&#39;-&#39;9&#39; 的字符串，编写一个算法来判断给定输入是否是累加数。

说明: 累加序列里的数不会以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。

示例 1:

输入: "112358"
输出: true
解释: 累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8


示例 2:

输入: "199100199"
输出: true
解释: 累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199

进阶:<br>
你如何处理一个溢出的过大的整数输入?

*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "1124"
	res := isAdditiveNumber(str)
	fmt.Println("res:", res)
}

type num struct {
	StartIndex int
	EndIndex   int
	Value      []byte
}

func newNum(startIndex int) *num {
	return &num{
		StartIndex: startIndex,
		EndIndex:   startIndex - 1,
		Value:      make([]byte, 0),
	}
}

func (n *num) Add(char byte) *num {
	n.EndIndex = n.EndIndex + 1
	n.Value = append(n.Value, char)
	return n
}

func (n *num) GenerateInt() int {
	n1, _ := strconv.Atoi(string(n.Value))
	return n1
}

func isAdditiveNumber(num string) bool {
	total := len(num)
	var i, j, k int
	var n1, n2, n3 int
	var exist = false
	temp1 := newNum(i)
	if total-2 == 0 {
		return true
	}
	for i = 0; i < total-2; i++ {
		temp1.Add(num[i])
		n1 = temp1.GenerateInt()
		temp2 := newNum(i + 1)
		for j = i + 1; j < total-1; j++ {
			temp2.Add(num[j])
			n2 = temp2.GenerateInt()
			temp3 := newNum(j + 1)
			for k = j + 1; k < total; k++ {
				temp3.Add(num[k])
				n3 = temp3.GenerateInt()
				if n3 == n1+n2 {
					exist = isAdditiveNumber(num[i+1:])
					return exist
				}
			}
		}
	}
	return exist
}
