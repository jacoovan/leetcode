/*
LCP 2:分式化简
leetcodeID : 100092
leetcode地址 : https://leetcode-cn.com/problems/deep-dark-fraction/
难度 : 简单

有一个同学在学习分式。他需要将一个连分数化成最简分数，你能帮助他吗？

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/09/09/fraction_example_1.jpg" style="height: 195px; width: 480px;">

连分数是形如上图的分式。在本题中，所有系数都是大于等于0的整数。

 

输入的cont代表连分数的系数（cont[0]代表上图的a<sub>0</sub>，以此类推）。返回一个长度为2的数组[n, m]，使得连分数的值等于n / m，且n, m最大公约数为1。

 

示例 1：

输入：cont = [3, 2, 0, 2]
输出：[13, 4]
解释：原连分数等价于3 + (1 / (2 + (1 / (0 + 1 / 2))))。注意[26, 8], [-13, -4]都不是正确答案。

示例 2：

输入：cont = [0, 0, 3]
输出：[3, 1]
解释：如果答案是整数，令分母为1即可。

限制：


	cont[i] >= 0
	1 <= cont的长度 <= 10
	cont最后一个元素不等于0
	答案的n, m的取值都能被32位int整型存下（即不超过2 ^ 31 - 1）。


 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func fraction(cont []int) []int {
    
}