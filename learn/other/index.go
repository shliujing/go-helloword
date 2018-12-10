package main

import "fmt"

//1. 指针需要初始化 分配地址
//2. 传指针值过去，修改值，而不是修改指针地址

func fn(r *string) {
	*r = "aaa"
}

func addr(r *string) {
	str := "bb"
	r = &str
}

func main() {
	//var a = new(string)
	var a *string
	a = new(string)
	*a = "abc"
	fmt.Println(a)//打印地址
	fmt.Println(*a)//打印指针值
	fn(a) // 传值修改是可以的
	fmt.Println(*a)

	// 传地址修改是不行的
	addr(a) // 传地址，修改地址指向是不行的
	fmt.Println(a)
}
