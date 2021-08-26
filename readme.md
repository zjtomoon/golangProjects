# go语言学习入门

## 线上go语言学习资源

[地鼠文档 (topgoer.cn)](http://www.topgoer.cn/)

[前景 · Go语言中文文档 (topgoer.com)](http://topgoer.com/)

[首页 - Go语言中文网 - Golang中文社区 (studygolang.com)](https://studygolang.com/)

## go语言面试题整理

### 1、使用值接收者实现接口和使用指针接收者实现接口的区别？
+ 1、 使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存
+ 2、 指针接收者实现接口只能存结构体指针类型的变量
### 2、关于数据库索引结构，B树和B+树的区别
+ B+ Tree(B Tree变种)
```
 1、非叶子节点不存储data,只存储索引(冗余),可以存放更多的索引
 2、叶子节点包含所有索引字段
 3、叶子节点用指针连接，提高区间访问的性能
```

![](https://markdown-pngs.oss-cn-shanghai.aliyuncs.com/go%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0/B%2BTree.png)
![](https://markdown-pngs.oss-cn-shanghai.aliyuncs.com/go%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0/B%2BTree2.jpg)

+ B Tree
```
 1、叶子节点具有相同的深度，叶子节点的指针为空
 2、所有索引元素不重复
 3、节点中的数据索引从左到右递增排列
```

![](https://markdown-pngs.oss-cn-shanghai.aliyuncs.com/go%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0/BTree.png)

### 3、Docker为什么比VM快？
+ 1、Docker有着比虚拟机更少的抽象层
+ 2、Docker利用的是宿主机的内核，VM需要Guest OS,所以说，新建一个容器的时候，docker不需要像虚拟机一样重新加载一个操作系统内核，避免引导。虚拟机加载GuestOs是分钟级别的，而Docker是利用宿主机的操作系统，是秒级的。

### 4、golang什么时候方法使用指针接收者？
+ 1、需要修改结构体变量的值时要使用指针接收者
+ 2、结构体本身比较大，拷贝的内存开销比较大时，也要使用指针接收者
+ 3、保持一致性:如果有一个方法使用了指针接收者，其他方法为了统一也要使用指针接收者。

### 5、自定义类型和类型别名的区别？

```go
 type MyInt int //自定义类型
 type NewInt = int //类型别名
 /*
    类型别名只在代码编写过程中存在和有效，编译完之后就不存在了。
    内置的byte和rune都属于类型别名
 */
```

## 《GO语言核心编程》学习笔记

# Go语言进阶
+ 使用go重写redis 
+ 使用go重写sqlite
+ redis、sqlite项目完成后分离到新建的github仓库，并发布linux(ubuntu14、centos7)、windows(windows 7)release版本
+ 学习kubernetes 1.15.11源码 自己编写
+ k8s项目完成后分离到新建的github仓库，并发布linux(ubuntu14、centos7)、windows(windows 7)release版本



