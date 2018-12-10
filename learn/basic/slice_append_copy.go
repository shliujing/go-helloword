package main

import "fmt"

func main() {
	var numbers []int
	printSlice2(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice2(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice2(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2,3,4)
	printSlice2(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1,numbers)
	printSlice2(numbers1)


	//badasda
	a := []int{1,2,3}
	b := []int{5,7,9}
	b = append(b,a...)
	printSlice2(b)
}


func printSlice2(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}