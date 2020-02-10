/*
814:二叉树剪枝
leetcodeID : 832
leetcode地址 : https://leetcode-cn.com/problems/binary-tree-pruning/
难度 : 中等

给定二叉树根结点 root ，此外树的每个结点的值要么是 0，要么是 1。

返回移除了所有不包含 1 的子树的原二叉树。

( 节点 X 的子树为 X 本身，以及所有 X 的后代。)


示例1:
输入: [1,null,0,0,1]
输出: [1,null,0,null,1]
 
解释: 
只有红色节点满足条件&ldquo;所有不包含 1 的子树&rdquo;。
右图为返回的答案。

<img alt="" src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/06/1028_2.png" style="width:450px" />



示例2:
输入: [1,0,1,0,0,0,1]
输出: [1,null,1,null,1]


<img alt="" src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/06/1028_1.png" style="width:450px" />



示例3:
输入: [1,1,0,1,1,0,1,0]
输出: [1,1,0,1,1,null,1]


<img alt="" src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/05/1028.png" style="width:450px" />


说明: 

<ul>
	给定的二叉树最多有 100 个节点。
	每个节点的值只会为 0 或 1 。
</ul>

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
    
}