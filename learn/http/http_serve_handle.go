package main

import (
	"log"
	"net/http"
)

//在 main 函数中我们只用了 http.NewServeMux 函数来创建一个空的 ServeMux。
//然后我们使用 http.RedirectHandler 函数创建了一个新的处理器，这个处理器会对收到的所有请求，都执行307重定向操作到 http://www.baidu.com。
//接下来我们使用 ServeMux.Handle 函数将处理器注册到新创建的 ServeMux，所以它在 URL 路径/foo 上收到所有的请求都交给这个处理器。
//最后我们创建了一个新的服务器，并通过 http.ListenAndServe 函数监听所有进入的请求，通过传递刚才创建的 ServeMux来为请求去匹配对应处理器。
func main() {
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://www.baidu.com", 307)
	mux.Handle("/foo", rh)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}