/*
472:连接词
leetcodeID : 472
leetcode地址 : https://leetcode-cn.com/problems/concatenated-words/
难度 : 困难

给定一个不含重复单词的列表，编写一个程序，返回给定单词列表中所有的连接词。

连接词的定义为：一个字符串完全是由至少两个给定数组中的单词组成的。

示例:


输入: ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]

输出: ["catsdogcats","dogcatsdog","ratcatdogcat"]

解释: "catsdogcats"由"cats", "dog" 和 "cats"组成;
     "dogcatsdog"由"dog", "cats"和"dog"组成;
     "ratcatdogcat"由"rat", "cat", "dog"和"cat"组成。


说明:


	给定数组的元素总数不超过 10000。
	给定数组中元素的长度总和不超过 600000。
	所有输入字符串只包含小写字母。
	不需要考虑答案输出的顺序。


*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{
		"cat",
		"cats",
		"catsdogcats",
		"dog",
		"dogcatsdog",
		"hippopotamuses",
		"rat",
		"ratcatdogcat",
	}
	res := findAllConcatenatedWordsInADict(words)
	fmt.Println("res:", res)
}

func findAllConcatenatedWordsInADict(words []string) []string {
	wordMap := setWordsLengthMap(words)
	return findMultiWords(wordMap)
}

func setWordsLengthMap(words []string) map[int][]string {
	wordMap := make(map[int][]string)
	for _, word := range words {
		length := len(word)
		_, ok := wordMap[length]
		if !ok {
			wordMap[length] = make([]string, 0)
		}
		wordMap[length] = append(wordMap[length], word)
	}
	return wordMap
}

func findMultiWords(wordMap map[int][]string) []string {
	lengths := make([]int, 0)
	for length := range wordMap {
		lengths = append(lengths, length)
	}
	sort.Ints(lengths)
	minLength := lengths[0]

	multiWords := make([]string, 0)
	for _, length := range lengths {
		words := wordMap[length]
		for _, word := range words {
			isMultiWord := wordContainSubword(word, wordMap, minLength)
			if isMultiWord {
				multiWords = append(multiWords, word)
			}
		}
	}
	return multiWords
}

func wordContainSubword(word string, subWordMap map[int][]string, minLength int) bool {
	for wordLen := range subWordMap {
		if wordLen >= len(word) {
			continue
		}
		subwords := subWordMap[wordLen]
		for _, subword := range subwords {
			if strings.Contains(word, subword) {
				remainWord := strings.ReplaceAll(word, subword, "")
				if len(remainWord) == 0 {
					return true
				}
				if wordContainSubword(remainWord, subWordMap, minLength) {
					return true
				}
			}
		}
	}
	return false
}
