
# 学习笔记

学习自：http://www.runoob.com/go/
``
## 1-9基础语法在 main.go

```
// 6. 函数
/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
   /* 声明局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result 
}

/* 调用函数并返回最大值 */
ret = max(a, b)

fmt.Printf( "最大值是 : %d\n", ret )

// 函数返回多个值
func swap(x, y string) (string, string) {
   return y, x
}

a, b := swap("Mahesh", "Kumar")
fmt.Println(a, b)

// 7. 局部/全局/形式变量（形式参数会作为函数的局部变量来使用。重名变量名，优先值以局部为准）

// 8. 数组
// var variable_name [SIZE] variable_type，
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0} // 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
 balance[4] = 50.0 // 赋值

// 9. 指针
var a int= 20   /* 声明实际变量 */
var ip *int        /* 指向整数型，声明指针变量 */

ip = &a  /* 指针变量的存储地址 */
fmt.Printf("a 变量的地址是: %x\n", &a  )

/* 指针变量的存储地址 */
fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

/* 使用指针访问值 */
fmt.Printf("*ip 变量的值: %d\n", *ip )
}

if(ptr == nil)    /* ptr 是空指针 */

// 10.结构体 struct.go

// 11. 切片 slice_append_copy.go, slice_cut.go

// 12. range
Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对的 key 值。不包含实际作用。


// 13. Map map_create_use_ok.go, map_delete.go
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

// 14. 递归函数 recursion_fibonacci.go
递归，就是在运行的过程中调用自己。

// 15. 类型转换，见 main.go
类型转换用于将一种数据类型的变量转换为另外一种类型的变量。Go 语言类型转换基本格式如下：

type_name(expression) //type_name 为类型，expression 为表达式。

// 16. 接口 interface_method_struct.go
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。方法是需要实现的。

// 17. 错误处理 error.go
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
type error interface {
    Error() string
}

18. 其他
疑问
:= 具体是 定义变量+赋值？

```

## orm

/Users/jingliu/lj-local/code/go/src/go-helloword/orm/hello-orm.go

![](http://i.iamlj.com/18-11-30/91016011.jpg)

## json

https://blog.csdn.net/zxy_666/article/details/80173288

```
interface{}类型在json.Unmarshal时，会自动将JSON转换为对应的数据类型： 
JSON的boolean 转换为bool 
JSON的数值 转换为float64 
JSON的字符串 转换为string 
JSON的Array 转换为[]interface{} 
JSON的Object 转换为map[string]interface{}，也可以转为本身类型 
JSON的null 转换为nil
```

