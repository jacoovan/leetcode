/*
2:两数相加
leetcodeID : 2
leetcode地址 : https://leetcode-cn.com/problems/add-two-numbers/
难度 : 中等

给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807


 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
    list1 := ListNode {
        Val : 2,
        Next : nil,
    }
    list1.Next = &ListNode {
        Val : 4,
        Next : nil,
    }
    list1.Next.Next = &ListNode {
        Val : 3,
        Next : nil,
    }

    list2 := ListNode {
        Val : 5,
        Next : nil,
    }
    list2.Next = &ListNode {
        Val : 6,
        Next : nil,
    }
    list2.Next.Next = &ListNode {
        Val : 4,
        Next : nil,
    }

    res := addTwoNumbers(&list1, &list2)

    for {
        fmt.Println(res)
        if res.Next != nil {
            res = res.Next
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    ret := ListNode {
        Val : 0,
        Next : nil,
    }
    add := false

    current := &ret
    for {
        result := l1.Val + l2.Val
        if add == true {
            result += 1
        }
        if result < 9 {
            (*current).Val = result
            add = false
        } else {
            (*current).Val = result - 10
            add = true
        }


        if l1.Next == nil && l2.Next == nil {
            (*current).Next = nil
            break
        } else {
            (*current).Next = & ListNode {
                Val : 0,
                Next : nil,
            }
            current = (*current).Next
        }

        if l1.Next != nil {
            l1 = l1.Next
        }
        if l2.Next != nil {
            l2 = l2.Next
        }
    }

    return &ret
}