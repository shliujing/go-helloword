package basic

import (
	"fmt"
	"unsafe"
)

// 2. 变量
var a = "菜鸟教程"
var b string = "runoob.com"
var c bool

// 3. 常量

const e string = "abc"
const f string = "aaa"

const (
	Unknown = 0
	Female  = 1
	Male    = 2

	a1 = "abc"
	c1 = unsafe.Sizeof(a)
	// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	d1 = iota
	e1 = iota
)

func main() {
	//1. hello world
	fmt.Println("hello go");
	println(a, b, c)
	println(e, f, Unknown, d1)

	// 4. 判断
	var a4 int = 20
	var b4 int = 22
	if (a4<b4) {
		fmt.Printf("a4小于b4")
	}

	// 5. 判断
	//for true {} ,for key, value := range oldMap  ,  for(;;)
	for c5 := 0; c5 < 3; c5++ {
		fmt.Printf("c5 的值为: %d\n", c5)
	}

	// 15. 强制转换
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)
}
