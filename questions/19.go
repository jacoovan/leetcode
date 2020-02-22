/*
19:删除链表的倒数第N个节点
leetcodeID : 19
leetcode地址 : https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
难度 : 中等

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.


说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

 */
package main

import(
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main(){

    head := &ListNode{
        1,
        nil,
    }
    head.Next = &ListNode{
        2,
        nil,
    }
    head.Next.Next = &ListNode{
        3,
        nil,
    }
    head.Next.Next.Next = &ListNode{
        4,
        nil,
    }
    head.Next.Next.Next.Next = &ListNode{
        5,
        nil,
    }
    n := 2
    head = removeNthFromEnd(head, n)

    for {
        if head == nil {
            break
        }
        fmt.Println(head)
        head = head.Next
    }
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var prevNode *ListNode
    var deleteNode *ListNode
    var nextNode *ListNode

    index := 1
    res := head
    for {
        if head == nil {
            break
        }

        if index - n == -1 {
            prevNode = res
        }

        if index - n == 0 {
            deleteNode = res
        }

        if index - n == 1 {
            nextNode = res
        }

        if head.Next != nil {
           head = head.Next
           index++
           if prevNode != nil && index - n > -1 {
                prevNode = prevNode.Next
           }
            if deleteNode != nil && index - n > 0 {
                deleteNode = deleteNode.Next
            }
           if nextNode != nil && index - n > 1 {
              nextNode = nextNode.Next
           }
        } else {
           head = nil
        }
    }

    if deleteNode != nil {
        if nextNode == nil && prevNode != nil {
            res = prevNode
        } else if nextNode != nil {
            nextNode.Next = prevNode
        }
    }

    return res
}