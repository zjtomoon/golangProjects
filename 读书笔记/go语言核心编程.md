# 《Go语言核心编程》

## 第一章 基础知识day01

### 1.2 初识Go程序

+ 第一个Go程序

[hello](../golang-core/day01/hello/main.go)

### 1.3.2 标识符

### 1.6.1 指针
[pointer](../golang-core/day01/pointer/main.go)

### 1.6.2 数组

[array](../golang-core/day01/array/main.go)

### 1.6.3 切片

[slice01](../golang-core/day01/slice/slice01/main.go)

[slice02](../golang-core/day01/slice/slice02/main.go)

### 1.6.4 map

[map](../golang-core/day01/map/main.go)
map的遍历

[map02](../golang-core/day01/map02/main.go)

### 1.6.5 struct
struct有两种形式：一种是struct类型字面量，另一种是使用type声明的自定义struct类型。
实际使用struct字面量的场景不多，更多的时候是通过type自定义一个新的类型来实现的。
type是自定义类型的关键字，不但支持struct类型的创建，还支持任意其他子定义类型的创建。

[struct](../golang-core/day01/structTest/main.go)
### 1.7.2 switch 语句

[switch](../golang-core/day01/switchTest/main.go)
### 1.7.3 for语句

[forTest](../golang-core/day01/forTest/for.md)
## 第二章 函数day02
### 2.1.3 实参到形参的传递
Go函数实参到形参的传递永远是值拷贝，有时函数调用后实参指向的值发生了变化，那是因为参数传递的是指针的拷贝，
实参是一个指针变量，传递给形参的是这个指针变量的副本，二者指向同一地址，本质上参数传递仍然是值拷贝。

[parameter](../golang-core/day02/parameter/main.go)
### 2.1.4 不定参数
函数的不定参数有如下几个特点：
(1)所有的不定参数类型必须是相同的
(2)不定参数必须是函数的最后一个参数
(3)不定参数名在函数体内相当于切片，对切片的操作同样适合对不定参数的操作。

[uncertain_parameter](../golang-core/day02/uncertain_parameter/main.go)

### 2.2.1 函数签名

[funcSig](../golang-core/day02/funcSig/main.go)

函数类型和map、slice、一样，实际函数类型变量和函数名都可以当作指针变量，该指针变量指向函数代码的开始位置。
通常说函数类型变量是一种引用类型，未初始化的函数类型的变量默认值是nil。

Go语言中函数是"第一公民"(first class)。有名函数的函数名可以看作函数类型的常量，可以直接使用函数名调用函数，
也可以直接赋值给函数类型变量，后续通过该变量来调用该函数。
```go
    package main
    
    func sum(a,b int) int {
     return a + b
    }
    
    func main()  {
     sum(3,4)
     f := sum
     f(1,2)
    }
```
### 2.2.2 匿名函数
Go提供两种函数：有名函数和匿名函数。匿名函数可以看作函数字面量，所以直接使用函数类型变量的地方都可以由匿名函数代替。
匿名函数可以直接赋值给函数变量，可以当作实参，也可以作为返回值，还可以直接被调用。
[anonymous_function](../golang-core/day02/anonymous_function/main.go)

### defer
Go函数里提供了defer关键字，可以注册多个延迟调用，这些调用以`先进后出(FILO)`的顺序在函数返回前被执行。
defer常用语保证一些资源最终一定能够得到回收和释放。

[defer](../golang-core/day02/defer/main.go)

defer后面必须是函数或方法的调用，不能是语句，否则会报expression in defer must be function call错误。
defer函数的实参在注册时通过值拷贝传递进去。
defer语句必须先注册后才能执行，如果defer位于return之后，则defer因为没有注册，不会执行。
```go

    package main
    
    func main()  {
     defer func() {
     	println("first")
     }()
     a := 0
     println(a)
     return 
     
     defer func() {
     	println("second")
     }()
    }
```
主动调用os.Exit(int)退出进程时，defer将不再被执行，即使defer已经提前注册。

[defer02](../golang-core/day02/defer02/main.go)

defer的好处是可以在一定程度上避免资源泄露。。特别是有很多return语句
defer语句的位置不当,有可能导致panic,一般defer语句放在错误检查语句之后
defer也有明显的副作用:defer会推迟资源的释放,defer尽量不要放到循环语句里面
defer语句的位置不当，有可能导致panic,一般defer语句放在错误检查语句之后。
defer也有明显的副作用：defer会推迟资源的释放，defer尽量不要放到循环语句里面。
## 第三章 类型系统day03
## 第四章 接口 day04
## 第五章 并发 day05
## 第六章 反射 day06
## 第七章 语言陷阱 day07
### 7.5 值、指针和引用
#### 7.5.1 传值还是传引用
+ Go只有一种参数传递规则，那就是值拷贝，这种规则包括两种含义

+ (1)函数参数传递时使用地是值拷贝
+ (2)实例赋值给接口变量，接口对实例地引用是值拷贝

+ 有时在明明是值拷贝的地方，结果却修改了变量的内容，有以下两种情况:
+ (1)直接传递的是指针。指针传递同样是值拷贝，但指针和指针副本的值指向的地址是同一个地方，所以能修改实参值。
+ (2)参数是复合数据类型，这些复合数据类型内部有指针类型的元素，此时参数的值拷贝并不影响指针的指向。

+ Go复合类型中chan、map、slice、interface内部都是通过指针指向具体的数据，这些类型的变量在作为函数参数传递时，实际上相当于指针的副本。 
## 第八章 工程管理 

