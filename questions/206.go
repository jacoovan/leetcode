/*
206:反转链表
leetcodeID : 206
leetcode地址 : https://leetcode-cn.com/problems/reverse-linked-list/
难度 : 简单

反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

进阶:<br>
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

 */
package main

import(
    "fmt"
)

func main(){
    root := &ListNode{
        1,
        nil,
    }

    root.Next = &ListNode{
        2,
        nil,
    }

    root.Next.Next = &ListNode{
        3,
        nil,
    }

    root.Next.Next.Next = &ListNode{
        4,
        nil,
    }

    root.Next.Next.Next.Next = &ListNode{
        5,
        nil,
    }
    fmt.Println("before")
    fmt.Println(root)
    tmp := root
    for {
        if tmp.Next != nil {
            tmp = tmp.Next
            fmt.Println(tmp)
        } else {
            break
        }
    }

    fmt.Println("after")
    root = reverseList(root)
    fmt.Println(root)
    tmp = root
    for {
       if tmp.Next != nil {
           tmp = tmp.Next
           fmt.Println(tmp)
       } else {
           break
       }
    }
}

type ListNode struct {
    Val int
    Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    temp := head
    tail := &ListNode{}
    for {
        current := temp
        if current.Next == nil {
            tail = current
            break
        }
        temp = current.Next
    }
    if head.Next == nil {
        return head
    }
    reverseNode(head, head.Next)
    head.Next.Next = head
    head.Next = nil
    return tail
}

func reverseNode(prev *ListNode, current *ListNode) {
    if current.Next == nil {
        return
    }

    prev = current
    current = current.Next
    reverseNode(prev, current)

    next := current.Next
    current.Next = prev
    prev.Next = next
}
