/*
477:汉明距离总和
leetcodeID : 477
leetcode地址 : https://leetcode-cn.com/problems/total-hamming-distance/
难度 : 中等

两个整数的 <a href="https://baike.baidu.com/item/%e6%b1%89%e6%98%8e%e8%b7%9d%e7%a6%bb/475174?fr=aladdin">汉明距离</a> 指的是这两个数字的二进制数对应位不同的数量。

计算一个数组中，任意两个数之间汉明距离的总和。

示例:


输入: 4, 14, 2

输出: 6

解释: 在二进制表示中，4表示为0100，14表示为1110，2表示为0010。（这样表示是为了体现后四位之间关系）
所以答案为：
HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.


注意:


	数组中元素的范围为从 0到 10^9。
	数组的长度不超过 10^4。


*/
package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 14, 2}
	res := totalHammingDistance(nums)
	fmt.Println("res:", res)
}

func totalHammingDistance(nums []int) int {
	length := len(nums) - 1
	distance := 0
	for i := range nums[:length] {
		for j := i + 1; j <= length; j++ {
			distance = distance + getHammingDistance(nums[i], nums[j])
		}
	}
	return distance
}

func getHammingDistance(n1, n2 int) int {
	biggerCode := make([]int, 0)
	smallerCode := make([]int, 0)
	if n1 > n2 {
		biggerCode = getHammingCode(n1)
		smallerCode = getHammingCode(n2)
	} else {
		biggerCode = getHammingCode(n2)
		smallerCode = getHammingCode(n1)
	}

	fmt.Println(biggerCode, smallerCode)

	distance := 0
	smallerLength := len(smallerCode)
	for i := range biggerCode {
		if i < smallerLength {
			if smallerCode[i] != biggerCode[i] {
				distance = distance + 1
			}
			continue
		}
		if biggerCode[i] == 1 {
			distance = distance + 1
		}
	}
	return distance
}

func getHammingCode(num int) []int {
	res := make([]int, 0)

	if num%2 == 1 {
		res = append(res, 1)
		num = num - 1
		num = num / 2
		res = append(res, getHammingCode(num)...)
		return res
	}

	if num > 0 && num%2 == 0 {
		res = append(res, 0)
		num = num / 2
		res = append(res, getHammingCode(num)...)
		return res
	}
	return res
}
