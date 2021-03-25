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
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	res := deleteDuplicates(tree)
	if res == nil {
		fmt.Println("res:", res)
		return
	}

	temp := res
	index := 1
	for temp != nil {
		fmt.Printf("res%d:%v", index, temp)
		temp = temp.Next
		index++
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

	var tempHeadNode *ListNode
	var tempNode *ListNode

	for {
		if tempNode != nil {
			tempNode.Next = getFirstNode(tempNode.Next)
			if tempNode.Next == nil {
				break
			}
			tempNode = tempNode.Next
		}

		if tempHeadNode == nil {
			tempHeadNode = getFirstNode(head)
			if tempHeadNode == nil {
				break
			}
			tempNode = tempHeadNode
		}
	}
	head = tempHeadNode
	return head
}

func getFirstNode(node *ListNode) *ListNode {
	var currentNode *ListNode = node
	if currentNode == nil {
		return nil
	}
	var duplicate = false
	for {
		if currentNode.Next != nil {
			if currentNode.Next.Val == currentNode.Val {
				duplicate = true
				currentNode = currentNode.Next
				continue
			}

			if !duplicate {
				break
			}

			if currentNode.Next.Next == nil {
				duplicate = false
				currentNode = currentNode.Next
				break
			}

			if currentNode.Next.Next != nil {
				if currentNode.Next.Next.Val != currentNode.Next.Val {
					duplicate = false
					currentNode = currentNode.Next
					break
				}
				currentNode = currentNode.Next.Next
			}
		} else {
			break
		}
	}
	if !duplicate {
		return currentNode
	}
	return nil
}
