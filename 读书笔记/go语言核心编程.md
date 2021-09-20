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

### 2.3defer
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

### 闭包
闭包是由函数及其相关引用环境组合而成的实体，一般通过在匿名函数中引用外部函数的局部变量或包全局变量构成

闭包=函数+引用环境

如果函数返回的闭包引用了该函数的局部变量（ 参数或函数内部变量〉
（1 ）多次调用该函数，返回的多个闭包所引用的外部变量是多个副本，原因是每次调用函
数都会为局部变量分配内存
(2 ）用一个闭包函数多次，如果该闭包修改了其引用的外部变量，则每一次调用该闭包对
该外部变量都有影响，因为闭包函数共享外部引用。

[closure](../golang-core/day02/closure/main.go)

如果一个函数调用返回的闭包引用修改了全局变 ，则每次调用都会影响全局变量。
如果函数返回的闭包引用的是全局变量 ，则多次调用该函数返回的多个闭包引用的都是
同理，调用 个闭包多次引用的也是同一个 。此时如果闭包中修改 值的逻辑，
每次闭包调用都会影响全局变量 的值。使用闭包是为了减少全局变量，所以闭包引用全局
变量不是好的编程方式。

[closure02](../golang-core/day02/closure02/main.go)

[closure03](../golang-core/day02/closure03/main.go)

闭包最初的目的是减少全局变量，在函数调用的过程中隐式地传递共享变量，有其有用的
面；但是这种隐秘的共享变 的方式带来的坏处是不够直接，不够清晰，除非是非常有价值
的地方， 般不建议使用闭包
对象是附有行为的数据，而闭包是附有数据的行为， 类在定义时已经显式地集中定义了行
为，但是闭包中的数据没有显式地集中声明的地方，这种数据和行为精合的模型不是一种推荐
的编程模型，闭包仅仅是锦上添花 东西，不是不可缺少的。有关闭包的内部实现原 2.7
节。

### panic和recover

panic recover 的函数签名如下
panic(i nterface{})
revover()interface{} 
panic 两种情况， 种是程序主动调用 panic 函数，另一种是程序产生运行时错误，
由运行时检测井抛出
发生 panic 后， 程序会从调用 panic 函数位置或发生 panic 的地方立即返回，逐层向上执
行函数的 defer 语句，然后逐层打印函数调用堆枝，直到被 recover 捕获或运行到最外层函数而
退出。
panic 参数是 个空接口类型 interface ｛｝，所以任意类型的变量都可以传递给 panic （空接
口的详细介绍见 4.3 节） 。调用 panic 方法非常简单 panic(xxx ）。
panic 不但可以在函数正常流程中抛出，在 defer 逻辑里也可以再次调用 panic 或抛出 panic a 
defer 里面的 panic 能够被后续执行的 defer 捕获。
recover（）用来捕获 pan ic ，阻止 panic 继续向上传递 recover （）和 defer 一起使用 ，
但是 `recover()只有在 defer 后面的函数体内被直接调用才能捕获 panic 终止异常，否则返回 nil ，异常继续向外传递`

```go
    //这个会捕获失败
    defer recover()
    
    //这个会捕获失败
    defer fmt.Println(recover())
    
    //这个两层嵌套也会捕获也会失败
    
   defer func() {
   	func() {
   		println("defer inner")
   		recover()
   	}()
   }()
   
   //如下场景会捕获成功
   defer func() {
   	    println("defer innner")
   	    recover()
   }()
   
   func except()  {
    recover()
   }
   
   func test()  {
        defer except()
        panic("test panic")
   }
```

[PanicandRecover](../golang-core/day02/PanicandRecover/main.go)

包中 in it 函数引发的 panic 只能在 in it 函数中捕获，在 main 中无法被捕获，原因是 in it
数先于 main 执行，有关包相关的知识参见 .3 节。函数并不能捕获内部新启动的 goroutine 所抛
出的 panic 。例如:

[panicTest02](../golang-core/day02/panicTest02/main.go)

## 第三章 类型系统day03

### 3.1.1 命名类型和未命名类型

[TypeTest01](../golang-core/day03/TypeTest03/main.go)

### 3.1.3 类型相同和类型赋值
两个命名类型是否相同，参考如下:
(1) 两个命名类型相同的条件是两个类型声明的语句完全相同。
(2) 命名类型和未命名类型永远不相同。
(3) 两个未命名类型相同的条件是它们的类型声明字面量的结构相同，并且内部元素的类型相同。
(4) 通过类型别名语句声明的两个类型相同。

[typeTest02](../golang-core/day03/typeTest02/main.go)

### 3.1.4 类型强制转换

[TypeTest03](../golang-core/day03/TypeTest03/main.go)

[typeTest04](../golang-core/day03/typeTest04/main.go)

### 自定义类型方法
+ 自定义structleix
+ struct初始化
    + (1) 按照字段顺序进行初始化
    + (2) 按照指定字段名进行初始化(推荐使用)
    + (3) 使用new创建内置函数，字段默认初始化为其类型的零值，返回值是指向结构的指针。
        ```go
          p := new(Person)
        //这种方法不常用，一般使用struct都不会将所有字段初始化为零值。
        ```
    + (4) 一次初始化一个字段
    + (5) 使用构造函数进行初始化
+ 结构字段的特点
    结构的字段可以是任意的类型，基本类型、接口类型、指针类型、函数类型都可以作为 struct
    字段。结 字段的类型名必须唯 struct 字段类型可以是普通类型，也可以是指针。
+ 匿名字段
+ 自定义接口类型

### 3.2.2 方法

[methodTest01](../golang-core/day03/methodTest01/main.go)

###3.3.1 一般方法调用

### 3.3.2 方法值

[MethodValue](../golang-core/day03/MethodValue/main.go)

### 3.3.3方法表达式

[methodExpression](../golang-core/day03/methodExpression/main.go)

### 3.3.4 方法集

[MethodSet](../golang-core/day03/MethodSet/main.go)

### 3.3.5 值调用和表达式调用的方法集

[methodSet02](../golang-core/day03/methodSet02/main.go)

### 3.4.1 组合

+ 内嵌字段的初始化和访问
[methodSet03](../golang-core/day03/methodSet03/main.go)

[methodSet04](../golang-core/day03/methodSet04/main.go)

+ 内嵌字段的方法调用

[methodSet05](../golang-core/day03/methodSet05/main.go)

### 3.4.2 组合的方法集

[methodSet06](../golang-core/day03/methodSet06/main.go)

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

