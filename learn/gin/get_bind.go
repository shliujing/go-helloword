package main

import "log"
import "github.com/gin-gonic/gin"
//ShouldBindQuery 函数只绑定字符串(get)，不绑定post数据。


type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	if c.BindQuery(&person) == nil {
		log.Println("====== Only Bind Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}
