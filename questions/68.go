/*
68:文本左右对齐
leetcodeID : 68
leetcode地址 : https://leetcode-cn.com/problems/text-justification/
难度 : 困难

给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。

你应该使用&ldquo;贪心算法&rdquo;来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 &#39; &#39; 填充，使得每行恰好有 maxWidth 个字符。

要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

文本的最后一行应为左对齐，且单词之间不插入额外的空格。

说明:

<ul>
	单词是指由非空格字符组成的字符序列。
	每个单词的长度大于 0，小于等于 maxWidth。
	输入单词数组 words 至少包含一个单词。
</ul>

示例:

输入:
words = ["This", "is", "an", "example", "of", "text", "justification."]
maxWidth = 16
输出:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]


示例 2:

输入:
words = ["What","must","be","acknowledgment","shall","be"]
maxWidth = 16
输出:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]
解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
     因为最后一行应为左对齐，而不是左右两端对齐。       
     第二行同样为左对齐，这是因为这行只包含一个单词。


示例 3:

输入:
words = ["Science","is","what","we","understand","well","enough","to","explain",
         "to","a","computer.","Art","is","everything","else","we","do"]
maxWidth = 20
输出:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]


 */
package main

import (
    "fmt"
    "math"
)

func main(){
    words := []string{"This", "is", "an", "example", "of", "text", "justification."}
    maxWidth := 16
    res := fullJustify(words, maxWidth)
    for i, p := range res {
        fmt.Println(i, "p:", "|"+p+"|", len(p))
    }
}

func fullJustify(words []string, maxWidth int) []string {
    totalLength := 0
    for _, word := range words {
        totalLength += len(word)
    }

    respWords := [][]string{}
    tempWords  := []string{}
    tempLength := 0
    for _, word := range words {
        wordLength := len(word)
        if tempLength + wordLength > maxWidth {
            respWords = append(respWords, tempWords)
            tempWords  = []string{word}
            tempLength = len(word)
            continue
        }
        tempLength += wordLength + 1
        tempWords  = append(tempWords, word)
    }
    respWords = append(respWords, tempWords)

    resp := []string{}
    for _, words := range respWords {
        totalWordsLength := 0
        for _, word := range words {
            totalWordsLength += len(word)
        }
        minWhiteSpaces := ""
        firstExtraWhiteSpaces := ""
        if len(words) > 1 {
            var minWhiteSpacesLength int = int(math.Floor(float64((maxWidth - totalWordsLength) / (len(words) - 1))))
            var firstExtraWhiteSpacesLength int = maxWidth - totalWordsLength - minWhiteSpacesLength * len(words)
            for i := 0; i < minWhiteSpacesLength; i++ {
                minWhiteSpaces += " "
            }
            for i := 0; i < minWhiteSpacesLength + firstExtraWhiteSpacesLength; i++ {
                firstExtraWhiteSpaces += " "
            }
        } else {
            for i := 0; i < maxWidth - totalWordsLength; i++ {
                firstExtraWhiteSpaces += " "
            }
        }

        temp := ""
        for i := len(words) - 1; i >= 0; i-- {
            if i == 0 {
                temp = words[0] + firstExtraWhiteSpaces + temp
                resp = append(resp, temp)
                temp = ""
                continue
            }
            temp = minWhiteSpaces + words[i] + temp
        }
    }
    return resp
}