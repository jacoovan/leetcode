/*
1028:从先序遍历还原二叉树
leetcodeID : 1093
leetcode地址 : https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal/
难度 : 困难

我们从二叉树的根节点 root 开始进行深度优先搜索。

在遍历中的每个节点处，我们输出 D 条短划线（其中 D 是该节点的深度），然后输出该节点的值。（如果节点的深度为 D，则其直接子节点的深度为 D + 1。根节点的深度为 0）。

如果节点只有一个子节点，那么保证该子节点为左子节点。

给出遍历输出 S，还原树并返回其根节点 root。

 

示例 1：

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/04/12/recover-a-tree-from-preorder-traversal.png" style="height: 200px; width: 320px;">

输入："1-2--3--4-5--6--7"
输出：[1,2,5,3,4,6,7]


示例 2：

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/04/12/screen-shot-2019-04-10-at-114101-pm.png" style="height: 250px; width: 256px;">

输入："1-2--3---4-5--6---7"
输出：[1,2,5,3,null,6,null,4,null,7]


示例 3：

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/04/12/screen-shot-2019-04-10-at-114955-pm.png" style="height: 250px; width: 276px;">

输入："1-401--349---90--88"
输出：[1,401,null,349,88,90]


 

提示：

<ul>
	原始树中的节点数介于 1 和 1000 之间。
	每个节点的值介于 1 和 10 ^ 9 之间。
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
func recoverFromPreorder(S string) *TreeNode {
    
}