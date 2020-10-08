/*
82:删除排序链表中的重复元素 II
leetcodeID : 82
leetcode地址 : https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
难度 : 中等

给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5


示例 2:

输入: 1->1->1->2->3
输出: 2->3

*/
package main

import "fmt"

func main() {
	tree := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	res := deleteDuplicates(tree)
	temp := res
	for temp != nil {
		fmt.Println("temp:", temp)
		temp = temp.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	lastNode := &ListNode{}
	lastNode = nil
	firstNode := head
	secondNode := head.Next
	if secondNode == nil {
		return head
	}
	duplicate := false
	for {
		if secondNode == nil {
			break
		}
		if firstNode.Val != secondNode.Val {
			if duplicate {
				lastNode = secondNode
			} else {
				lastNode = firstNode
			}
			break
		}
		duplicate = true
		firstNode = secondNode
		secondNode = firstNode.Next
	}
	head = lastNode

	for {
		if secondNode == nil {
			break
		}
		if firstNode.Val != secondNode.Val {
			lastNode = firstNode
			firstNode = secondNode
			secondNode = firstNode.Next
			continue
		}
		for {
			if secondNode == nil {
				break
			}
			firstNode = secondNode
			secondNode = firstNode.Next
			if secondNode == nil {
				lastNode.Next = nil
				break
			}
			if firstNode.Val == secondNode.Val {
				firstNode = secondNode
				secondNode = firstNode.Next
				continue
			}
			if secondNode.Next == nil {
				lastNode.Next = secondNode
				break
			}
			firstNode = secondNode
			secondNode = firstNode.Next
		}
	}
	return head
}
