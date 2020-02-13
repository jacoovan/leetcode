/*
138:复制带随机指针的链表
leetcodeID : 138
leetcode地址 : https://leetcode-cn.com/problems/copy-list-with-random-pointer/
难度 : 中等

给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。

要求返回这个链表的 <a href="https://baike.baidu.com/item/深拷贝/22785317?fr=aladdin" target="_blank">深拷贝</a>。 

我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

<ul>
	val：一个表示 Node.val 的整数。
	random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
</ul>

 

示例 1：

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e1.png" style="height: 138px; width: 680px;">

输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]


示例 2：

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e2.png" style="height: 111px; width: 680px;">

输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]


示例 3：

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e3.png" style="height: 119px; width: 680px;">

输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]


示例 4：

输入：head = []
输出：[]
解释：给定的链表为空（空指针），因此返回 null。


 

提示：

<ul>
	-10000 <= Node.val <= 10000
	Node.random 为空（null）或指向链表中的节点。
	节点数目不超过 1000 。
</ul>

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

