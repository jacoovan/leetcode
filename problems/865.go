/*
865:具有所有最深结点的最小子树
leetcodeID : 896
leetcode地址 : https://leetcode-cn.com/problems/smallest-subtree-with-all-the-deepest-nodes/
难度 : 中等

给定一个根为 root 的二叉树，每个结点的深度是它到根的最短距离。

如果一个结点在整个树的任意结点之间具有最大的深度，则该结点是最深的。

一个结点的子树是该结点加上它的所有后代的集合。

返回能满足&ldquo;以该结点为根的子树中包含所有最深的结点&rdquo;这一条件的具有最大深度的结点。

 

示例：

输入：[3,5,1,6,2,0,8,null,null,7,4]
输出：[2,7,4]
解释：
<img alt="" src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/01/sketch1.png" style="height: 238px; width: 280px;">
我们返回值为 2 的结点，在图中用黄色标记。
在图中用蓝色标记的是树的最深的结点。
输入 "[3, 5, 1, 6, 2, 0, 8, null, null, 7, 4]" 是对给定的树的序列化表述。
输出 "[2, 7, 4]" 是对根结点的值为 2 的子树的序列化表述。
输入和输出都具有 TreeNode 类型。


 

提示：

<ul>
	树中结点的数量介于 1 和 500 之间。
	每个结点的值都是独一无二的。
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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    
}