package main

import "fmt"

func main() {

	a := []*string{}
	b := []*string{}

	var str = new(string)
	*str = "aa"
	b = append(*b, *str)

	a = append(a, b...)
	fmt.Println(a)
	a = append(a, b...)
	fmt.Println(a)
}

func r(i int,data []*string) {
	i--
	//b := []string{"a", "b"}
}
