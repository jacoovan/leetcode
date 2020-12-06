/*
223:矩形面积
leetcodeID : 223
leetcode地址 : https://leetcode-cn.com/problems/rectangle-area/
难度 : 中等

在二维平面上计算出两个由直线构成的矩形重叠后形成的总面积。

每个矩形由其左下顶点和右上顶点坐标表示，如图所示。

<img alt="rectangle area" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rectangle_area.png">

示例:

输入: -3, 0, 3, 4, 0, -1, 9, 2
输出: 45

说明: 假设矩形面积不会超出 int 的范围。

*/
package main

import (
	"fmt"
)

func main() {
	area := computeArea(-3, 0, 3, 4, 0, -1, 9, 2)
	fmt.Println("area:", area)
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	rect1 := generateRect(A, B, C, D)
	rect2 := generateRect(E, F, G, H)
	duplicateArea := rect1.CoverRectArea(rect2)
	if duplicateArea == 0 {
		duplicateArea = rect2.CoverRectArea(rect1)
	}
	rArea1 := rect1.GenerateArea()
	rArea2 := rect2.GenerateArea()

	return rArea1 + rArea2 - duplicateArea
}

func generateRect(a, b, c, d int) *rect {
	leftTop := &point{X: a, Y: d}
	leftBottom := &point{X: a, Y: b}
	rightTop := &point{X: c, Y: d}
	rightBottom := &point{X: c, Y: b}
	return newRect(leftTop, leftBottom, rightTop, rightBottom)
}

type rect struct {
	LeftTop     *point
	RightTop    *point
	LeftBottom  *point
	RightBottom *point
}

func (r *rect) GenerateArea() int {
	width := r.RightTop.X - r.LeftTop.X
	height := r.LeftTop.Y - r.LeftBottom.Y
	return width * height
}

func (r *rect) CoverRectArea(r1 *rect) int {
	if r1 == nil {
		return 0
	}
	lt := r1.pointInRect(r.LeftTop)
	lb := r1.pointInRect(r.LeftBottom)
	rt := r1.pointInRect(r.RightTop)
	rb := r1.pointInRect(r.RightBottom)
	if !lt && !lb && !rt && !rb {
		return 0
	}
	width := 0
	height := 0
	area := 0
	if lt && !rt {
		width = r1.RightTop.X - r.LeftTop.X
	}
	if lt && rt {
		width = r.RightTop.X - r.LeftTop.X
	}
	if !lt && rt {
		width = r.RightTop.X - r1.LeftTop.X
	}

	if lb && !rb {
		width = r1.RightBottom.X - r.LeftBottom.X
	}
	if lb && rb {
		width = r.RightBottom.X - r.LeftBottom.X
	}
	if !lb && rb {
		width = r.RightBottom.X - r1.LeftBottom.X
	}

	if lt && !lb {
		height = r.LeftTop.Y - r1.LeftBottom.Y
	}
	if lt && lb {
		height = r.LeftTop.Y - r.LeftBottom.Y
	}
	if !lt && lb {
		height = r1.LeftTop.Y - r.LeftBottom.Y
	}

	if rt && !rb {
		height = r.RightTop.Y - r1.RightBottom.Y
	}
	if rt && rb {
		height = r.RightTop.Y - r.RightBottom.Y
	}
	if !rt && rb {
		height = r1.RightTop.Y - r.RightBottom.Y
	}

	area = width * height
	return area
}

func (r *rect) pointInRect(p *point) bool {
	if p.X >= r.RightTop.X || p.X <= r.LeftTop.X {
		return false
	}
	if p.Y >= r.RightTop.Y || p.Y <= r.RightBottom.Y {
		return false
	}
	return true
}

type point struct {
	X int
	Y int
}

func newRect(leftTop, leftBottom, rightTop, rightBottom *point) *rect {
	return &rect{
		LeftTop:     leftTop,
		RightTop:    rightTop,
		LeftBottom:  leftBottom,
		RightBottom: rightBottom,
	}
}
