/*
715:Range 模块
leetcodeID : 715
leetcode地址 : https://leetcode-cn.com/problems/range-module/
难度 : 困难

Range 模块是跟踪数字范围的模块。你的任务是以一种有效的方式设计和实现以下接口。

<ul>
	addRange(int left, int right) 添加半开区间 [left, right)，跟踪该区间中的每个实数。添加与当前跟踪的数字部分重叠的区间时，应当添加在区间 [left, right) 中尚未跟踪的任何数字到该区间中。
	queryRange(int left, int right) 只有在当前正在跟踪区间 [left, right) 中的每一个实数时，才返回 true。
	removeRange(int left, int right) 停止跟踪区间 [left, right) 中当前正在跟踪的每个实数。
</ul>

 

示例：

addRange(10, 20): null
removeRange(14, 16): null
queryRange(10, 14): true （区间 [10, 14) 中的每个数都正在被跟踪）
queryRange(13, 15): false （未跟踪区间 [13, 15) 中像 14, 14.03, 14.17 这样的数字）
queryRange(16, 17): true （尽管执行了删除操作，区间 [16, 17) 中的数字 16 仍然会被跟踪）


 

提示：

<ul>
	半开区间 [left, right) 表示所有满足 left <= x < right 的实数。
	对 addRange, queryRange, removeRange 的所有调用中 0 < left < right < 10^9。
	在单个测试用例中，对 addRange 的调用总数不超过 1000 次。
	在单个测试用例中，对  queryRange 的调用总数不超过 5000 次。
	在单个测试用例中，对 removeRange 的调用总数不超过 1000 次。
</ul>

 

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

type RangeModule struct {
    
}


func Constructor() RangeModule {
    
}


func (this *RangeModule) AddRange(left int, right int)  {
    
}


func (this *RangeModule) QueryRange(left int, right int) bool {
    
}


func (this *RangeModule) RemoveRange(left int, right int)  {
    
}


/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */