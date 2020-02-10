/*
1206:设计跳表
leetcodeID : 1337
leetcode地址 : https://leetcode-cn.com/problems/design-skiplist/
难度 : 困难

不使用任何库函数，设计一个跳表。

跳表是在 O(log(n)) 时间内完成增加、删除、搜索操作的数据结构。跳表相比于树堆与红黑树，其功能与性能相当，并且跳表的代码长度相较下更短，其设计思想与链表相似。

例如，一个跳表包含 [30, 40, 50, 60, 70, 90]，然后增加 80、45 到跳表中，以下图的方式操作：

<img alt="" src="https://assets.leetcode.com/uploads/2019/09/27/1506_skiplist.gif" style="width: 500px;"><br>
<small>Artyom Kalinin [CC BY-SA 3.0], via <a href="https://commons.wikimedia.org/wiki/file:skip_list_add_element-en.gif" target="_blank" title="artyom kalinin [cc by-sa 3.0 (https://creativecommons.org/licenses/by-sa/3.0)], via wikimedia commons">Wikimedia Commons</a></small>

跳表中有很多层，每一层是一个短的链表。在第一层的作用下，增加、删除和搜索操作的时间复杂度不超过 O(n)。跳表的每一个操作的平均时间复杂度是 O(log(n))，空间复杂度是 O(n)。

在本题中，你的设计应该要包含这些函数：

<ul>
	bool search(int target) : 返回target是否存在于跳表中。
	void add(int num): 插入一个元素到跳表。
	bool erase(int num): 在跳表中删除一个值，如果 num 不存在，直接返回false. 如果存在多个 num ，删除其中任意一个即可。
</ul>

了解更多 : <a href="https://en.wikipedia.org/wiki/skip_list" target="_blank">https://en.wikipedia.org/wiki/Skip_list</a>

注意，跳表中可能存在多个相同的值，你的代码需要处理这种情况。

样例:

Skiplist skiplist = new Skiplist();

skiplist.add(1);
skiplist.add(2);
skiplist.add(3);
skiplist.search(0);   // 返回 false
skiplist.add(4);
skiplist.search(1);   // 返回 true
skiplist.erase(0);    // 返回 false，0 不在跳表中
skiplist.erase(1);    // 返回 true
skiplist.search(1);   // 返回 false，1 已被擦除


约束条件:

<ul>
	0 <= num, target <= 20000
	最多调用 50000 次 search, add, 以及 erase操作。
</ul>

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

type Skiplist struct {
    
}


func Constructor() Skiplist {
    
}


func (this *Skiplist) Search(target int) bool {
    
}


func (this *Skiplist) Add(num int)  {
    
}


func (this *Skiplist) Erase(num int) bool {
    
}


/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */