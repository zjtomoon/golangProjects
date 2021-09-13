# 《Go web编程》要点整理

## 第三章 接受请求

+ Go语言拥有一系列成熟的标准库，如net/http和html/template，这些标准库可以用于构建web应用。
+ Go语言的net/http标准库可以将Http通信放在SSL之上进行，也就是通过https方式创建出更为安全的通信连接。
+ Go语言的处理器可以是任何带有ServeHTTP方法的结构，其中ServeHTTP方法需要接收两个参数：第一个参数是一个ResponseWriter接口，而第二个参数则是一个指向Request结构的指针。
+ 处理器函数是与处理器拥有相似行为的函数。处理器函数用于处理请求，它们和ServeHTTP方法拥有相同的签名。
+ 通过串联处理器或者处理器函数，可以对程序中的横切关注点进行分隔，并以模块化的方式处理请求。
+ 多路复用器也是处理器。比如，ServeMux就是一个HTTP请求多路复用器，它接受HTTP请求并根据请求中的URL将请求重定向到正确的处理器。DefaultServeMux是ServeMux的一个公开的实例，这个实例会被用作默认的多路复用器。
+ Go 1.6或以上的版本中，net/http标准库默认支持HTTP/2。

## 第四章 处理请求

+ Request结构的Form、PostForm和MultipartFrom字段可以让用户更容易地提取出请求中的不同数据：用户只要调用ParseForm方法或者ParseMultipartForm方法对请求进行语法分析，然后访问相应的字段，就可以取得请求中包含的数据。
+ Form字段存储的是来自URL以及HTML表单的URL编码数据，Post字段存储的是来自HTML表单的URL编码数据，而MultipartForm字段存储的则是来自URL以及HTML表单的multipart编码数据。
+ 服务器通过向ResponseWriter写入首部和主体来向客户端返回响应。
+ 通过向ResponseWriter写入cookie,服务器可以将数据持久地存储在客户端上。
+ cookie可以用于实现闪现消息。

## 第五章

+ 在Web应用中，模板引擎会把模板和数据进行合并，生成将要返回给客户端地HTML。
+ Go地标准模板定义在html/template包当中。
+ Go模板引擎的工作方式就是对一个模板进行语法分析，接着在执行这个模板的时候，将一个ResponseWriter以及一些数据传递给它。被调用的模板引擎会对传入的已分析模板以及数据进行合并，然后把合并的结果传递给ResponseWriter。

