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

## 第六章

+ 通过使用结构将数据存储在内存里面，以此来构建数据缓存机制并提高响应速度。
+ 通过使用csv或者gob二进制格式将数据存储在文件里面，可以对用户提交的文件进行处理，或者为缓存数据提供备份。
+ 通过使用database/sql包，可以对关系数据库执行CRUD操作，并在不同的数据之间建立起相应的关系。
+ 通过sqlx和orm这样的第三方数据访问库，可以使用威力更强大的工具去操纵数据库中的数据。

## 第七章

+ Web服务主要分为两种类型——一种是基于SOAO的Web服务，而另一种则是基于REST的Web服务。
  + SAOP是一种协议，能够对定义在XML中的结构化数据进行交换。但是，因为SAOP的WSDL报文有可能变得非常复杂，所以基于SAOP的Web服务没有基于REST的Web服务那么流行。
  + 基于REST的Web服务通过HTTP协议向外界公开自己拥有的资源，并允许外界通过HTTP协议对这些资源执行指定的动作。
+ 创建何分析XML以及JSON的步骤都是相似的，用户要么根据指定的结构去生成XML或JSON,要么从指定的结构里面提取数据到XML或者JSON里面，前一种操作称为封装，而后一种操作则成为解封。

