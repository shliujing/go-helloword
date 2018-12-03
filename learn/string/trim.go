package main

import (
	"log"
	"strings"
	"fmt"
)

// refer: https://blog.csdn.net/chenbaoke/article/details/40318423

func main() {
	log.Println(strings.TrimRight("abba", "ba"))
	log.Println(strings.TrimRight("abcdaaaaa", "abcd")) //匹配后，包括自己+右边，都删除掉
	log.Println(strings.TrimSuffix("abcddcba", "dcba"))

	//func Trim(s string, cutset string) string这个函数大家估计就看的多了，这个函数是去除两边的自定义字符的
	fmt.Println(strings.Trim("&&&&&nihao&&&&", "&")) //nihao

	//func TrimFunc(s string, f func(rune) bool) string该函数大家是不是看的多了，对就是自定义函数清楚，那我们自定义函数完成上边那个函数的效果吧
	fmt.Println(strings.TrimFunc("&&&&nihao&&&&", func(r rune) bool {
		if r == '&' {
			return true
		}
		return false
	})) //nihao

	//func TrimLeft(s string, cutset string) string该函数是删除左边的定义的字符
	fmt.Println(strings.TrimLeft("&widuu&", "&")) //widuu&

	//func TrimSpace(s string) string,清楚文本里边的空白操作\r\n\t
	fmt.Println(strings.TrimSpace("\r\t\n widuu \t")) //widuu

	//func TrimSuffix(s, suffix string) string这个函数去除后缀的，跟TrimPrefix这个正好相反，我们看实例
	fmt.Println(strings.TrimSuffix("web_widuu", "_widuu")) //web

	fmt.Println(strings.TrimSuffix("Hello 世界!!!!!", "!!!!")) // "Hello 世界!"
	//func TrimRight(s string, cutset string) string // TrimRight 将删除 s 尾部连续的包含在 cutset 中的字符
	fmt.Println(strings.TrimRight(" Hello 世界! ", " 世界!")) // " Hello"

	//注：TrimSuffix只是去掉s字符串结尾的suffix字符串，只是去掉１次，而TrimRight是一直去掉s字符串右边的字符串，只要有响应的字符串就去掉，是一个多次的过程，这也是二者的本质区别．
}
