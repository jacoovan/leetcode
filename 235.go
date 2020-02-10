/*
235:二叉搜索树的最近公共祖先
leetcodeID : 235
leetcode地址 : https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
难度 : 简单

给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

<a href="https://baike.baidu.com/item/%e6%9c%80%e8%bf%91%e5%85%ac%e5%85%b1%e7%a5%96%e5%85%88/8918834?fr=aladdin" target="_blank">百度百科</a>中最近公共祖先的定义为：&ldquo;对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。&rdquo;

例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]

<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/binarysearchtree_improved.png" style="height: 190px; width: 200px;">

 

示例 1:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6 
解释: 节点 2 和节点 8 的最近公共祖先是 6。


示例 2:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。

 

说明:

<ul>
	所有节点的值都是唯一的。
	p、q 为不同节点且均存在于给定的二叉搜索树中。
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
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  
}