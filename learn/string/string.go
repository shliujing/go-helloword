package main

import (
	"fmt"
	"strings"
)

func main()  {

	data:=[4]byte{1,2,3,4}
	var strs string
	//for i := range data {
	//	strings.Join(strs, (string)i)
	//}

	strings.Replace(strings.Trim(fmt.Sprint(data), "[]"), " ", ",", -1)

	fmt.Println(strings.Replace(strings.Trim(fmt.Sprint(data), "[]"), " ", ",", -1))
	fmt.Println(strs)
}
