/*
92:反转链表 II
leetcodeID : 92
leetcode地址 : https://leetcode-cn.com/problems/reverse-linked-list-ii/
难度 : 中等

反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:<br>
1 &le; m &le; n &le; 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
    list := &ListNode{
        Val:  1,
        Next: nil,
    }
    list.Next = &ListNode{Val: 2, Next: nil}
    list.Next.Next = &ListNode{Val: 3, Next: nil}
    list.Next.Next.Next = &ListNode{Val: 4, Next: nil}
    list.Next.Next.Next.Next = &ListNode{Val: 5, Next: nil}

    result := reverseBetween(list, 1, 4)
    _ = result
    temp := result
    for temp.Next != nil {
      fmt.Println("val:", temp.Val)
      temp = temp.Next
    }
    fmt.Println("val:", temp.Val)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    var i = 1
    nextNode := head.Next

    prevNode := head
    if m == 1 {
        prevNode = nil
    }
    endNode := &ListNode{Val: head.Val, Next: head.Next}
    for i < m {
        if i > 1 {
            prevNode = prevNode.Next
        }
        endNode = nextNode
        nextNode = nextNode.Next
        i++
    }

    startNode := endNode
    for i >= m && i < n {
        currentNode := nextNode
        nextNode = nextNode.Next
        tempNode := startNode
        startNode = currentNode
        startNode.Next = tempNode
        i++
    }
    afterNode := nextNode
    endNode.Next = afterNode

    if prevNode == nil {
        return startNode
    }
    prevNode.Next = startNode
    endNode.Next = afterNode
    return head
}