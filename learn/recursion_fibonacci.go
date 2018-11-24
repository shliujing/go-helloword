package learn

import "fmt"

// 递归，就是在运行的过程中调用自己。

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main(){
	var i int
	for i = 0; i< 10; i++{
		fmt.Printf("%d\t ",fibonacci(i))
	}
}
