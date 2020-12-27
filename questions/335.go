/*
335:路径交叉
leetcodeID : 335
leetcode地址 : https://leetcode-cn.com/problems/self-crossing/
难度 : 困难

给定一个含有 n 个正数的数组 x。从点 (0,0) 开始，先向北移动 x[0] 米，然后向西移动 x[1] 米，向南移动 x[2] 米，向东移动 x[3] 米，持续移动。也就是说，每次移动后你的方位会发生逆时针变化。

编写一个 O(1) 空间复杂度的一趟扫描算法，判断你所经过的路径是否相交。



示例 1:

┌───┐
│   │
└───┼──>
    │

输入: [2,1,1,2]
输出: true


示例 2:

┌──────┐
│      │
│
│
└────────────>

输入: [1,2,3,4]
输出: false


示例 3:

┌───┐
│   │
└───┼>

输入: [1,1,1,1]
输出: true


*/
package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 1, 4}
	res := isSelfCrossing(x)
	fmt.Println("res:", res)
}

func isSelfCrossing(x []int) bool {
	if len(x) < 4 {
		return false
	}

	if x[0] >= x[2] && x[1] <= x[3] {
		return true
	}

	if len(x) == 4 {
		return false
	}

	lastTop := x[0]
	lastLeft := -x[1]
	lastBottom := x[0] - x[2]
	lastRight := x[3] - x[1]

	inner := x[0] > x[2] || x[1] > x[3]
	x = x[4:]
	for i := range x {
		if inner {
			if i%4 == 0 && lastBottom+x[i] >= lastTop {
				lastTop = lastBottom + x[i]
				return true
			}
			if i%4 == 1 && lastRight-x[i] <= lastLeft {
				lastLeft = lastRight - x[i]
				return true
			}
			if i%4 == 2 && lastTop-x[i] <= lastBottom {
				lastBottom = lastTop - x[i]
				return true
			}
			if i%4 == 3 && lastLeft+x[i] >= lastRight {
				lastRight = lastLeft + x[i]
				return true
			}
			continue
		}
		if i%4 == 0 && lastBottom+x[i] < lastTop {
			lastTop = lastBottom + x[i]
			if i >= 4 {
				lastRight = lastLeft + (x[i-1] - x[i-3])
				if x[i] >= x[i-2]-x[i-4] && x[i] <= x[i-2] {
					lastLeft = lastRight
				}
			}
			inner = true
			continue
		}

		if i%4 == 1 && lastRight-x[i] > lastLeft {
			lastLeft = lastRight - x[i]
			if i >= 4 {
				lastTop = lastBottom + (x[i-1] - x[i-3])
				if x[i] >= x[i-2]-x[i-4] && x[i] <= x[i-2] {
					lastBottom = lastTop
				}
			}
			inner = true
			continue
		}
		if i%4 == 2 && lastTop-x[i] > lastBottom {
			lastBottom = lastTop - x[i]
			if i >= 4 {
				lastLeft = lastRight - (x[i-1] - x[i-3])
				if x[i] >= x[i-2]-x[i-4] && x[i] <= x[i-2] {
					lastRight = lastLeft
				}
			}
			inner = true
			continue
		}
		if i%4 == 3 && lastLeft+x[i] < lastRight {
			lastRight = lastLeft + x[i]
			if i >= 4 {
				lastBottom = lastTop - (x[i-1] - x[i-3])
				if x[i] >= x[i-2]-x[i-4] && x[i] <= x[i-2] {
					lastTop = lastBottom
				}
			}
			inner = true
			continue
		}
	}
	return false
}
