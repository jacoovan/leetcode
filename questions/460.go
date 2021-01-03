/*
460:LFU缓存
leetcodeID : 460
leetcode地址 : https://leetcode-cn.com/problems/lfu-cache/
难度 : 困难

设计并实现<a href="https://baike.baidu.com/item/%e7%bc%93%e5%ad%98%e7%ae%97%e6%b3%95">最不经常使用（LFU）</a>缓存的数据结构。它应该支持以下操作：get 和 put。

get(key) - 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。<br />
put(key, value) - 如果键不存在，请设置或插入值。当缓存达到其容量时，它应该在插入新项目之前，使最不经常使用的项目无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，最近最少使用的键将被去除。

进阶：<br />
你是否可以在 O(1) 时间复杂度内执行两项操作？

示例：

*/

// LFUCache cache = new LFUCache( 2 /* capacity (缓存容量) */ );

/*
cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回 1
cache.put(3, 3);    // 去除 key 2
cache.get(2);       // 返回 -1 (未找到key 2)
cache.get(3);       // 返回 3
cache.put(4, 4);    // 去除 key 1
cache.get(1);       // 返回 -1 (未找到 key 1)
cache.get(3);       // 返回 3
cache.get(4);       // 返回 4

*/
package main

import "fmt"

func main() {
	cache := Constructor(2)
	cache.Get(1)
	cache.Put(1, 1)
	cache.Get(1)
	cache.Put(2, 2)
	cache.Get(1)
	cache.Get(2)
	cache.Put(3, 3)
	cache.Get(1)
	cache.Get(3)
}

type LFUCache struct {
	len       int
	actualLen int
	lastNode  *cacheNode
	node      *cacheNode
	nodeMap   map[interface{}]*cacheNode
}

type cacheNode struct {
	prevNode *cacheNode
	nextNode *cacheNode
	key      interface{}
	value    interface{}
	count    int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		len:       capacity,
		actualLen: 0,
		nodeMap:   make(map[interface{}]*cacheNode),
	}
}

func (this *LFUCache) Print(operation string) {
	node := this.node
	i := 1
	for {
		if node == nil {
			break
		}
		fmt.Print(operation, "----index----", i)
		fmt.Print("----node----", node.key, node.value, node.count)
		node = node.nextNode
		i = i + 1
	}
	fmt.Println("")
}

func (this *LFUCache) Get(key int) int {
	defer this.Print(fmt.Sprintf("GET(%d)", key))
	v, ok := this.nodeMap[key]
	if !ok {
		return -1
	}
	v.count = v.count + 1
	this.sortNode(v)
	value, ok := v.value.(int)
	if !ok {
		return 0
	}
	return value
}

func (this *LFUCache) Put(key int, value int) {
	defer this.Print(fmt.Sprintf("PUT(%d)", key))
	v, ok := this.nodeMap[key]
	if !ok {
		actualLen := this.actualLen
		if this.actualLen < this.len {
			actualLen = actualLen + 1
		} else {
			delete(this.nodeMap, this.lastNode.key)
			fmt.Println("----actuallen(delete)----", this.lastNode.key, this.lastNode)
			if this.lastNode.prevNode != nil {
				this.lastNode.prevNode.nextNode = nil
				this.lastNode = this.lastNode.prevNode
			}
		}
		fmt.Println("----actuallen----", actualLen, key)
		node := &cacheNode{
			prevNode: nil,
			nextNode: this.node,
			key:      key,
			value:    value,
			count:    1,
		}
		if this.node != nil {
			this.node.prevNode = node
		}
		if this.actualLen == 0 {
			this.lastNode = node
		}
		this.nodeMap[key] = node
		this.node = node
		this.actualLen = actualLen
		return
	}

	v.count = v.count + 1
	this.sortNode(v)
}

func (this *LFUCache) sortNode(v *cacheNode) {
	if v.prevNode != nil && v.prevNode.count < v.count {
		if this.lastNode == v {
			this.lastNode = v.prevNode
		}
		if v.prevNode.prevNode == nil {
			this.node = v
		}
		tempPrevNode := v.prevNode.prevNode
		tempNextNode := v.nextNode
		v.nextNode = v.prevNode
		v.prevNode.prevNode = v
		v.prevNode.nextNode = tempNextNode
		v.prevNode = tempPrevNode
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
