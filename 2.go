// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
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

    ret := ListNode {
        Val : 0,
        Next : nil,
    }
    add := false

    current := &ret
    for {
        result := list1.Val + list2.Val
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


        if list1.Next == nil && list2.Next == nil {
            (*current).Next = nil
            break
        } else {
            (*current).Next = & ListNode {
                Val : 0,
                Next : nil,
            }
            current = (*current).Next
        }

        if list1.Next != nil {
            list1 = *list1.Next
        }
        if list2.Next != nil {
            list2 = *list2.Next
        }
    }

    for {
        fmt.Print(ret.Val)
        if ret.Next == nil {
            break
        }
        ret = *ret.Next
    }
}