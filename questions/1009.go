/*
1009:十进制整数的反码
leetcodeID : 1054
leetcode地址 : https://leetcode-cn.com/problems/complement-of-base-10-integer/
难度 : 简单

每个非负整数 N 都有其二进制表示。例如， 5 可以被表示为二进制 "101"，11 可以用二进制 "1011" 表示，依此类推。注意，除 N = 0 外，任何二进制表示中都不含前导零。

二进制的反码表示是将每个 1 改为 0 且每个 0 变为 1。例如，二进制数 "101" 的二进制反码为 "010"。

给定十进制数 N，返回其二进制表示的反码所对应的十进制整数。

 




示例 1：

输入：5
输出：2
解释：5 的二进制表示为 "101"，其二进制反码为 "010"，也就是十进制中的 2 。


示例 2：

输入：7
输出：0
解释：7 的二进制表示为 "111"，其二进制反码为 "000"，也就是十进制中的 0 。


示例 3：

输入：10
输出：5
解释：10 的二进制表示为 "1010"，其二进制反码为 "0101"，也就是十进制中的 5 。


 

提示：


	0 <= N < 10^9


 */
package main

import(
    "fmt"
)

func main(){
    n := 10

    res := bitwiseComplement(n)

    fmt.Println(res)
}

func bitwiseComplement(N int) int {
    res := genBinary(N)

    reverseBinary := []int {}
    for _, num := range res {
        temp := 0
        if num == 0 {
            temp = 1
        }
        reverseBinary = append(reverseBinary, temp)
    }

    num := 0
    temp := 1
    for i := len(reverseBinary) - 1; i >= 0; i-- {
        num += temp * reverseBinary[i]
        temp = 2 * temp
    }

    return num
}

func genBinary(n int) []int {
	if n == 0 {
		return []int {0}
	}

	binary := []int {}
	for {
			if n == 1 {
				binary = append(binary, 1)
				break
			}
			tmp := n % 2
			binary = append(binary, tmp)
			n = n - tmp

			n = int(n / 2)
	}

	res := []int {}
	for i := len(binary) - 1; i >= 0; i = i-1 {
		res = append(res, binary[i])
	}

	return res
}