
## go http

实现一个最简单HTTP server需要多少代码？只需要一行，Python2的python -m SimpleHTTPServer，ruby的ruby -run -e httpd . -p 8888。对于Golang，实现一个最简单的http server也用不着几行，却能带来更具杀伤力的性能。


参考：https://studygolang.com/articles/9467

一切的基础：ServeMux 和 Handler
Go 语言中处理 HTTP 请求主要跟两个东西相关：ServeMux 和 Handler。

ServrMux 本质上是一个 HTTP 请求路由器（或者叫多路复用器，Multiplexor）。它把收到的请求与一组预先定义的 URL 路径列表做对比，然后在匹配到路径的时候调用关联的处理器（Handler）。

处理器（Handler）负责输出HTTP响应的头和正文。任何满足了http.Handler接口的对象都可作为一个处理器。通俗的说，对象只要有个如下签名的ServeHTTP方法即可：

ServeHTTP(http.ResponseWriter, *http.Request)