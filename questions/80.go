/*
80:删除排序数组中的重复项 II
leetcodeID : 80
leetcode地址 : https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/
难度 : 中等

给定一个排序数组，你需要在<a href="http://baike.baidu.com/item/%e5%8e%9f%e5%9c%b0%e7%ae%97%e6%b3%95" target="_blank">原地</a>删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在<a href="https://baike.baidu.com/item/%e5%8e%9f%e5%9c%b0%e7%ae%97%e6%b3%95" target="_blank">原地</a>修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定 nums = [1,1,1,2,2,3],

函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。

你不需要考虑数组中超出新长度后面的元素。

示例 2:

给定 nums = [0,0,1,1,1,1,2,3,3],

函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。

你不需要考虑数组中超出新长度后面的元素。


说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以&ldquo;引用&rdquo;方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

// nums 是以&ldquo;引用&rdquo;方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);
gb
// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}

*/
package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	// nums := []int{1, 1, 1, 2, 2, 3}
	res := removeDuplicates(nums)
	fmt.Println("res:", res)
}

func removeDuplicates(nums []int) int {
	var flag = -1
	var times = 0
	var length = 0
	for i, num := range nums {
		if num != flag {
			flag = num
			times = 1
			continue
		}
		times++

		if times == 3 {
			times = 1
			pos := 0
			for j := i; j < len(nums); j++ {
				if nums[j] > flag {
					pos = j
					break
				}
			}
			var temp int
			for j := pos; j < len(nums); j++ {
				temp = nums[j]
				nums[j] = nums[i+(j-pos)]
				nums[i+(j-pos)] = temp
			}
			flag = nums[i]
			continue
		}
	}
	var prev = nums[0]
	for _, num := range nums {
		if prev <= num {
			prev = num
			length++
		}
	}
	fmt.Println("nums:", nums)
	return length
}
