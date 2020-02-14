/*
863:二叉树中所有距离为 K 的结点
leetcodeID : 893
leetcode地址 : https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree/
难度 : 中等

给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 K 。

返回到目标结点 target 距离为 K 的所有结点的值的列表。 答案可以以任何顺序返回。

 




示例 1：

输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2

输出：[7,4,1]

解释：
所求结点为与目标结点（值为 5）距离为 2 的结点，
值分别为 7，4，以及 1

<img alt="" src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/28/sketch0.png" style="height: 240px; width: 280px;">

注意，输入的 "root" 和 "target" 实际上是树上的结点。
上面的输入仅仅是对这些对象进行了序列化描述。


 

提示：


	给定的树是非空的，且最多有 K 个结点。
	树上的每个结点都具有唯一的值 0 <= node.val <= 500 。
	目标结点 target 是树上的结点。
	0 <= K <= 1000.


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
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
    
}