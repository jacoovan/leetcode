/*
1106:解析布尔表达式
leetcodeID : 1197
leetcode地址 : https://leetcode-cn.com/problems/parsing-a-boolean-expression/
难度 : 困难

给你一个以字符串形式表述的 <a href="https://baike.baidu.com/item/%e5%b8%83%e5%b0%94%e8%a1%a8%e8%be%be%e5%bc%8f/1574380?fr=aladdin" target="_blank">布尔表达式</a>（boolean） expression，返回该式的运算结果。

有效的表达式需遵循以下约定：

<ul>
	"t"，运算结果为 True
	"f"，运算结果为 False
	"!(expr)"，运算过程为对内部表达式 expr 进行逻辑 非的运算（NOT）
	"&(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 与的运算（AND）
	"|(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 或的运算（OR）
</ul>

 

示例 1：

输入：expression = "!(f)"
输出：true


示例 2：

输入：expression = "|(f,t)"
输出：true


示例 3：

输入：expression = "&(t,f)"
输出：false


示例 4：

输入：expression = "|(&(t,f,t),!(t))"
输出：false


 

提示：

<ul>
	1 <= expression.length <= 20000
	expression[i] 由 {&#39;(&#39;, &#39;)&#39;, &#39;&&#39;, &#39;|&#39;, &#39;!&#39;, &#39;t&#39;, &#39;f&#39;, &#39;,&#39;} 中的字符组成。
	expression 是以上述形式给出的有效表达式，表示一个布尔值。
</ul>

 */
package main

import(
    "fmt"
)

func main(){
    fmt.Println("请完成你的逻辑代码")
}

func parseBoolExpr(expression string) bool {
    
}