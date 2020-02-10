/*
382:链表随机节点
leetcodeID : 382
leetcode地址 : https://leetcode-cn.com/problems/linked-list-random-node/
难度 : 中等

给定一个单链表，随机选择链表的一个节点，并返回相应的节点值。保证每个节点被选的概率一样。

进阶:<br />
如果链表十分大且长度未知，如何解决这个问题？你能否使用常数级空间复杂度实现？

示例:


// 初始化一个单链表 [1,2,3].
ListNode head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
Solution solution = new Solution(head);

// getRandom()方法应随机返回1,2,3中的一个，保证每个元素被返回的概率相等。
solution.getRandom();


 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
    
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */