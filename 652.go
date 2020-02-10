/*
652:寻找重复的子树
leetcodeID : 652
leetcode地址 : https://leetcode-cn.com/problems/find-duplicate-subtrees/
难度 : 中等

给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。

两棵树重复是指它们具有相同的结构以及相同的结点值。

示例 1：

        1
       / \
      2   3
     /   / \
    4   2   4
       /
      4


下面是两个重复的子树：

      2
     /
    4


和

    4


因此，你需要以列表的形式返回上述重复子树的根结点。

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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    
}