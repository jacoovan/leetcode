/*
83:删除排序链表中的重复元素
leetcodeID : 83
leetcode地址 : https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
难度 : 简单

给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2


示例 2:

输入: 1->1->2->3->3
输出: 1->2->3

*/
package main

import (
	"fmt"
)

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	list = deleteDuplicates(list)
	fmt.Println("val:", list.Val)
	for list.Next != nil {
		list = list.Next
		fmt.Println("val:", list.Val)
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
	var prevNode *ListNode = nil
	var currentNode *ListNode = head
	var nextNode *ListNode = head.Next

	_ = prevNode

	for nextNode != nil {
		currentValue := currentNode.Val
		nextValue := nextNode.Val

		if nextValue != currentValue {
			prevNode = currentNode
			currentNode = nextNode
			nextNode = nextNode.Next
			continue
		}

		prevNode = currentNode
		currentNode.Next = nextNode.Next
		nextNode = nextNode.Next
	}
	return head
}
