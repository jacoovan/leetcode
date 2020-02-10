/*
352:将数据流变为多个不相交区间
leetcodeID : 352
leetcode地址 : https://leetcode-cn.com/problems/data-stream-as-disjoint-intervals/
难度 : 困难

给定一个非负整数的数据流输入 a<sub>1</sub>，a<sub>2</sub>，&hellip;，a<sub>n，</sub>&hellip;，将到目前为止看到的数字总结为不相交的区间列表。

例如，假设数据流中的整数为 1，3，7，2，6，&hellip;，每次的总结为：

[1, 1]
[1, 1], [3, 3]
[1, 1], [3, 3], [7, 7]
[1, 3], [7, 7]
[1, 3], [6, 7]


 

进阶：<br>
如果有很多合并，并且与数据流的大小相比，不相交区间的数量很小，该怎么办?

提示：<br>
特别感谢 <a href="https://discuss.leetcode.com/user/yunhong">@yunhong</a> 提供了本问题和其测试用例。

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

type SummaryRanges struct {
    
}


/** Initialize your data structure here. */
func Constructor() SummaryRanges {
    
}


func (this *SummaryRanges) AddNum(val int)  {
    
}


func (this *SummaryRanges) GetIntervals() [][]int {
    
}


/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */