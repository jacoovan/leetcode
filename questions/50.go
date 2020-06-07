/*
50:Pow(x, n)
leetcodeID : 50
leetcode地址 : https://leetcode-cn.com/problems/powx-n/
难度 : 中等

实现 <a href="https://www.cplusplus.com/reference/valarray/pow/" target="_blank">pow(x, n)</a> ，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000


示例 2:

输入: 2.10000, 3
输出: 9.26100


示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2<sup>-2</sup> = 1/2<sup>2</sup> = 1/4 = 0.25

说明:

<ul>
	-100.0 < x < 100.0
	n 是 32 位有符号整数，其数值范围是 [&minus;2<sup>31</sup>, 2<sup>31 </sup>&minus; 1] 。
</ul>

 */
package main

import(
    "fmt"
)

func main(){
    x := 2.0
    n := -2
    res := myPow(x, n)
    fmt.Println(res)
}

func myPow(x float64, n int) float64 {
    temp := x

    var n1 int
    if n > 0 {
    	n1 = n
    } else {
    	n1 = n * -1
    }

    if n1 == 0 {
    	return 1
    } else if n1 == 1 {
    	if n < 0 {
	    	return 1 / temp
	    }
    	return x
    }
    for n1 > 1 {
    	i := n1 % 2
    	if i == 1 {
    		temp = float64(i) * temp
    		n1 = n1 - 1
    	}

    	switch i {
    	case 1:
    		temp = float64(i) * temp
    		n1 = n1 - 1
    	default :
    		n1 = int(float64(n1) / 2)
    		temp = temp * temp
    	}
    }

    if n < 0 {
    	return 1 / temp
    }

    return temp
}