# 《Go web编程》要点整理

## 第三章 接受请求

+ Go语言拥有一系列成熟的标准库，如net/http和html/template，这些标准库可以用于构建web应用。
+ Go语言的net/http标准库可以将Http通信放在SSL之上进行，也就是通过https方式创建出更为安全的通信连接。
+ Go语言的处理器可以是任何带有ServeHTTP方法的结构，其中ServeHTTP方法需要接收两个参数：第一个参数是一个ResponseWriter接口，而第二个参数则是一个指向Request结构的指针。
+ 处理器函数是与处理器拥有相似行为的函数。处理器函数用于处理请求，它们和ServeHTTP方法拥有相同的签名。
+ 通过串联处理器或者处理器函数，可以对程序中的横切关注点进行分隔，并以模块化的方式处理请求。
+ 多路复用器也是处理器。比如，ServeMux就是一个HTTP请求多路复用器，它接受HTTP请求并根据请求中的URL将请求重定向到正确的处理器。DefaultServeMux是ServeMux的一个公开的实例，这个实例会被用作默认的多路复用器。
+ Go 1.6或以上的版本中，net/http标准库默认支持HTTP/2。

