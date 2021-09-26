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

### 函数类型

[funcDecaration](../golang-core/day03/funcDecaration/main.go)

## 第四章 接口 day04
### 4.1.1接口声明

### 4.1.2接口初始化
+ 实例赋值接口
+ 接口变量赋值接口变量

### 4.1.3接口方法调用

[interfaceTest01](../golang-core/day04/interfaceTest01/main.go)

### 4.2.1 类型断言

[interfaceTypeAssertion](../golang-core/day04/interfaceTypeAssertion/main.go)

[interfaceTest01](../golang-core/day04/interfaceTest02/main.go)

### 4.2.2类型查询

[TypeSwitches](../golang-core/day04/TypeSwitches/main.go)

[TypeSwitches02](../golang-core/day04/TypeSwitches02/main.go)

### 4.2.3接口优点和使用形式
+ 接口优点
(1)解耦：复杂系统进行垂 和水平的分割是常用的设计手段，在层与层之间使用接口进
行抽象和解辑是 种好的编程策略 Go 的非侵入式的接口使层与层之间的代码更加干净，
类型和实现的接口之间不需要显式声明，增加了接口使用的自由度
(2)实现泛型：由于现阶段 Go 语言还不支持泛型，使用空接口作为函数或方法参数能够
用在需要泛型的场景中
+ 接口使用形式
(1) 作为结构内嵌字段。
(2) 作为函数或方法的形参。
(3) 作为函数或方法的返回值。
(4) 作为其他接口定义的嵌入字段。

### 4.3.2 空接口的用途

+ 空接口和泛型
Go语言没有泛型，如果一个函数需要接收任意类型的参数，则参数类型可以使用空接口类型，这是弥补没有泛型的一种手段。
+ 空接口与反射
空接口是反射实现的基础，反射库就是将相关具体的类型转换并赋值给空接口后才去处理。

### 4.3.3 空接口和nil

[null-Interface](../golang-core/day04/null-Interface/main.go)

### 4.4.1 接口内部实现-数据结构

非空接口的底层数据结构是iface
```go
    //src/runtime/runtime2.go
    type iface struct {
    	tab  *itab //itab存放类型及方法指针信息
    	data unsafe.Pointer //数据信息
    }
```
可以看到iface结构，有两个指针类型字段
+ itab:用来存放接口自身类型和绑定的实例类型及实例相关的函数指针
+ 数据指针data:指向接口绑定的实例的副本，接口的初始化也是一种值拷贝

itab数据结构
```go
    //src/runtime/runtime2.go
    type itab struct {
    	inter *interfacetype //接口自身的静态类型
    	_type *_type //_type就是接口存放的具体实例的类型(动态类型)
    	hash  uint32 //hash存放具体类型的Hash值
    	_     [4]byte
    	fun   [1]uintpr
    }
```
itab有5个字段
+ inner是指向接口类型元信息的指针
+ _type是指向接口存放的具体类型元信息的指针，iface里的data指针指向的是该类型的值。一个是类型信息，另一个是类型的值。
+ hash是具体类型的Hash值,_type里面也有hash,这里冗余存放主要是为了接口断言或类型查询时快速访问
+ fun是一个函数指针

Go语言类型元信息最初由编译器负责构建，并以表的形式存放在编译后的对象文件中，再由链接器在链接时进行段合并、符号重定向。
这些类型信息在接口的动态调用和反射中被运行时引用。
接口的类型元信息的数据结构：

```go
    //描述接口的类型
    type interfacetype struct {
    	typ     _type //类型通用部分
    	pkgpath name //接口所属包的名字信息，name内存放的不仅有名称，还有描述信息
    	mhdr    []imethd //接口的方法
    }
    //接口方法元信息
    type imethod struct {
    	name nameOff //方法名在编译后的section里面的偏移量
    	ityp typeOff //方法类型在编译后的section里面的偏移量
    }
```
### 4.4.4 空接口数据结构
```go
    //src/runtime/runtime2.go
    
    //空接口
    type eface struct {
    	_type *_type
    	data unsafe.Pointer
    }
```
从eface的数据结构可以看出，空接口不是真的为空，其保留了具体实例的类型和值拷贝，即便存放的具体类型是空的，
空接口也不是空的。
由于空接口自身没有方法集，所以空接口变量实例化后的真正用途不是接口方法的动态调用。空接口在go语言中真正的意义是
`支持多态`。有如下几种方式使用了空接口(将空接口类型还原)
(1) 通过接口类型断言
(2) 通过接口类型查询
(3) 通过反射

## 第五章 并发 day05
### 5.1.1 并发和并行
并发和并行是两个不同的概念：
+ 并行意味着程序在任意时刻都是同时运行的
+ 并发意味着程序在单位时间内是同时运行的
### 5.1.2 goroutine
+ 通过go + 匿名函数形式启动goroutine
[goroutineTest01](../golang-core/day05/goroutineTest01/main.go)

+ 通过go + 有名函数形式启动goroutine
[goroutineTest02](../golang-core/day05/goroutineTest02/main.go)

goroutine有如下特性
+ go 的执行是非阻塞的，不会等待。
+ go后面的函数返回值会被忽略。
+ 调度器不能保证多个goroutine的执行次序。
+ 没有父子goroutine的概念，所有的goroutine是平等地调度和执行的。
+ Go程序在执行时会单独为main函数创建一个goroutine,遇到其他go关键字时再去创建其他的goroutine。
+ Go没有暴露goroutine id给用户，所以不能在一个goroutine里面显式地操作另一个goroutine,不过runetime包
提供了一些访问和设置goroutine的相关信息。

1.func GoMAXPROCS 用来设置或查询可以并发执行的goroutine数目，n大于1表示设置GOMAXPROCS值，否则表示当前的GOMAXPROCS值

[goMaxProcs](../golang-core/day05/goMaxProcs/main.go)

2.func Goexit
func Goexit()是结束当前goroutine的运行，Goexit在结束当前goroutine运行之前会调用当前goroutine已经注册的defer。
Goexit并不会产生panic。所以该goroutine defer里面的recover调用都返回nil。

3.func Gosched
func Gosched()是放弃当前调度执行机会，将当前goroutine放到队列中等待下次被调度。

###5.1.3 chan

Go的哲学是“不要通过共享内存来通信，而是通过通信来共享内存”。通道是Go通过通信来共享内存的载体。

[channelTest01](../golang-core/day05/channelTest01/main.go)

goroutine运行结束后退出，写到缓冲通道中的数据不会消失，它可以缓冲和适配两个goroutine处理速率不一致的情况，
缓冲通道和消息队列类似，有削峰和增大吞吐量的功能。

[channelTest02](../golang-core/day05/channelTest02/main.go)

操作不同状态的chan会引发三种行为。
panic
+ (1)向已经关闭的通道写数据会导致panic
     + 最佳实践是由写入者关闭通道，能最大程度地避免向已经关闭的通道写数据而导致的panic  
+ (2)重复关闭的通道会导致panic

阻塞
+ (1)向未初始化的通道写数据或读数据都会导致当前goroutine的永久阻塞
+ (2)向缓冲区已满的通道写入数据会导致goroutine阻塞
+ (3)通道中没有数据，读取该通道会导致

非阻塞
+ (1)读取已经关闭的通道不会引发阻塞，而是立即返回通道元素类型的零值，可以使用comma,ok语法判断通道是否关闭
+ (2)向有缓冲且没有满的通道读/写不会引发阻塞

### 5.1.4 waitGroup
WaitGroup用来等待多个goroutine完成，main goroutine调用Add设置需要等待goroutine的数目，
每一个goroutine结束时调用Done(),Wait()被main用来等待所有的goroutine完成。

[waitGroup](../golang-core/day05/waitGroup/main.go)

### 5.1.5 select

[select](../golang-core/day05/select/main.go)

### 5.1.6 扇入(Fan in)和扇出(Fan out)

### 5.1.7 通知退出机制
读取已经关闭的通道不会引起阻塞，也不会导致panic,而是立即返回该通道存储类型的零值。
关闭select监听的某个通道能使select立即感知这种通知，然后进行相应的处理，这就是所谓的`退出通知机制`。

[CloseChanneltoBroadcast](../golang-core/day05/CloseChanneltoBroadcast/main.go)

### 5.2.1 生成器
+ (1) 最简单的带缓冲的生成器，例如：

[generatorTest01](../golang-core/day05/generatorTest01/main.go)

+ (2) 多个goroutine增强型生成器，例如：

[generatorTest02](../golang-core/day05/generatorTest02/main.go)

+ (3) 有时希望生成器能够自动退出，可以借助Go通道的退出机制实现。例如：

[generatorTest03](../golang-core/day05/generatorTest03/main.go)

+ (4) 一个融合了并发、缓冲、退出通知等多重特性的生成器。例如：

[generatorTest04](../golang-core/day05/generatorTest04/main.go)

### 5.2.2 管道
通道可以分为两个方向，一个是读，一个是写，假如一个函数的输入参数和输出参数是相同的chan类型，则该函数可以调用自己，
最终形成一个调用链。

[pipeTest](../golang-core/day05/pipeTest/main.go)

### 5.2.3 每个请求一个goroutine

以计算100个自然数的和来举例，将计算任务拆分为多个task,每个task启动一个goroutine进行处理，程序实例代码如下

[goRoutineTest03](../golang-core/day05/goRoutineTest03/main.go)

程序的逻辑分析：
+ (1) InitTask函数构建task并发送到task通道中。
+ (2) 分发任务函数DistributeTask为每个task启动一个goroutine处理任务，等待其处理完成，然后关闭结果通道。
+ (3) ProcessResult函数读取并统计所有的结果。

这几个函数分别在不同的goroutine中运行，它们通过通道和sync.WaitGroup进行通信和同步。

### 5.2.4 固定worker工作池

服务器编程中使用最多的就是通过线程池来提升服务的并发处理能力。在Go语言编程中，一样可以轻松地构建固定数目的
goroutine作为工作线程池。下面以计算多个整数的和为例来说明这种并发范式。

程序除了主要的main goroutine,还开启了如下几类goroutine:
+ (1) 初始化任务的goroutine
+ (2) 分发任务的goroutine
+ (3) 等待所有worker结束通知，然后关闭结果通道的goroutine。

程序采用三个通道，分别是：
+ (1) 传递task任务的通道
+ (2) 传递task结果的通道
+ (3) 接收worker处理完任务后所发送通知的通道。

[workPool](../golang-core/day05/workerPool/main.go)

程序的逻辑分析：

+ (1) 构建task并发送到task通道中
+ (2) 分别启动n个工作线程，不停地从task通道中获取任务，然后将结果写入结果通道。如果任务通道被关闭，则负责向收敛结果的goroutine发送通知，告诉其当前worker已经完成工作。
+ (3) 收敛结果的goroutine接收到所有task已经处理完毕的信号后，主动关闭结果通道。
+ (4) main中的函数ProcessResult读取并统计所有的结果。

### 5.2.5 future模式

编程中经常遇到在一个流程中需要调用多个子调用的情况，这些子调用相互之间没有依赖，如果串行地调用，则耗时会很长，此时可以使用Go并发编程中的future模式。

future模式的基本工作原理:

+ (1) 使用chan作为函数参数。
+ (2) 启动goroutine调用函数。
+ (3) 通过chan传入参数。
+ (4) 做其他可以并行处理的事情。
+ (5) 通过chan异步获取结果。

[futureMod](../golang-core/day05/futureMod/main.go)

future最大的好处是将函数的同步调用转换为异步调用，适用于一个交易需要多个子调用且这些子调用没有依赖的场景。

### 5.3.1 context的设计目的
context库的设计目的就是跟踪goroutine调用树，并在这些goroutine调用树种传递通知和元数据。两个目的：
+ (1) 退出通知机制——通知可以传递给整个goroutine调用树上得每一个goroutine。
+ (2) 传递数据——数据可以传递给整个goroutine调用树上地每一个goroutine

###5.3.2 context基本数据结构

context包的整体工作机制:

第一个创建Context的goroutine被称为root节点。root节点负责创建一个实现Context接口的具体对象，并将该对象作为
参数传递到其新拉起的goroutine,下游的goroutine可以继续封装该对象，再传递到更下游的goroutine。Context对象
在传递的过程中最终形成一个树状的数据结构，这样通过位于root节点(树的根节点)的Context对象就能遍历整个Context对象树，
通知和消息就可以通过root节点传递出去，实现了上游goroutine对下游goroutine的消息传递。

+ Context接口

Context是一个基本接口，所有的Context对象都要实现该接口，context的使用者在调用接口中都使用了Context作为参数类型。
+ canceler接口

canceler接口是一个扩展接口，规定了取消通知的Context具体类型需要实现的接口，context包中的具体类型*cancelCtx和*timerCtx都实现了该接口

+ empty Context接口

emptyCtx 实现了Context接口，但不具备任何功能，因为其所有的方法都是空实现。其存在的目的是作为Context对象树的根(root节点)。
因为context包的使用思路就是不停地调用context包提供的包装函数来创建具有特殊功能的Context实例，每一个Context实例
的创建都以上一个Context对象为参数，最终形成一个树状的结构。

+ cancelCtx

cancelCtx是一个实现了Context接口的具体类型，同时实现了canceler接口。cancel具有退出通知方法。

+ timerCtx

timerCtx是一个实现了Context接口的具体类型，内部封装了cancelCtx类型实例，同时有一个deadline变量，用来实现定时退出通知。

+ valueCtx

valueCtx是一个实现了Context接口的具体类型，内部封装了Context接口类型，同时封装了一个k/v的存储变量。
valueCtx可用来传递通知信息。

[contextLib01](../golang-core/day05/contextLib01/Context.go)

### 5.3.3 API函数

With包装函数用来构建不同功能的Context具体对象。

+ (1) 创建一个带有退出通知的Context具体对象，内部创建一个cancelCtx的类型实例。

```go
func WithCancel(parent Context) (ctx Context,cancel CancelFunc)
```

+ (2) 创建一个带有超时通知的Context具体对象，内部创建一个timeCtx的类型实例。

```go
    func WithDeadline(parent Context,deadline time.Time) (Context,CancelFunc)
```

+ (3) 创建一个带有超时通知的Context具体对象，内部创建一个timeCtx的类型实例。

```go
    func WithTimeout(parent Context,timeout time.Duration) (Context,CancelFunc)
```

+ (4) 创建一个能够传送数据的Context具体对象，内部创建一个valueCtx的类型实例。

```go
    func WithValue(parent Context,key,val interface{}) Context
```

### 5.3.5 context的用法

[contextLib02](../golang-core/day05/contextLib02/main.go)

使用Context包的一般流程：
+ (1) 创建一个Context根函数
    ```go
      func Background() Context
      func TODO() Context
    ```
+ (2) 包装上一步创建的Context对象，使其具有特定的功能。

这些包函数是context package的核心，几乎所有的封装都是从包装函数开始的。原因是，使用Context包的核心就是使用期退出通知广播功能。

+ (3) 将上一步创建的对象作为实参传给后续启动的并发函数(通常作为函数的第一个参数),每个并发函数内部可以继续使用包装函数对传进来的
Context对象进行包装，添加自己所需的功能。

+ (4) 顶端的goroutine在超时后调用cancel退出通知函数，通知后端的所有goroutine释放资源。

+ (5) 后端的goroutine通过select监听Context.Done()返回的chan，及时响应前端goroutine的退出通知，一般停止本次处理，释放所占用的资源。

### 5.3.6 使用context传递数据的争议

在context中传递数据的坏处：

+ (1) 传递的都是interface{}类型的值，编译器不能通过进行严格的类型检验。
+ (2) 从interface{}到具体类型需要使用类型断言和接口查询，有一定的运行期开销和性能损失。
+ (3) 值在传递过程中有可能被后续的服务覆盖，且不易被发现。
+ (4) 传递信息不简明，较晦涩；不能通过代码或文档一眼看到传递的是什么，不利于后续维护。

context应该传递什么数据
+ (1) 日志信息
+ (2) 调试信息
+ (3) 不影响业务主逻辑的可选数据

context包提供的核心的功能是多个goroutine之间的退出通知机制，传递数据只是一个辅助功能，应该谨慎使用context传递数据。

### 5.4.2 调度模型

调度原则：
+ (1) 尽可能让每个cpu核心都有事情做
+ (2) 尽可能提高每个cpu核心做事情的效率

应用程序并发模型：
+ (1) 多进程模型

进程都被多核cpu并发调度，优点是每个进程都有自己独立的内存空间，隔离性好、健壮性高；缺点是进程比较重，进程的切换消耗较大，
进程间的通信需要多次在内核区和用户区之间复制数据。

+ (2) 多线程模型

这里的多线程是指启动多个内核线程进行处理，线程的优点是通过共享内存进行通信更快捷，切换代价小；缺点是多个线程共享内存空间，
极易导致数据访问混乱，某个线程误操作内存挂掉可能会危及整个线程组，健壮性不高。

+ (3) 用户级多线程模型

用户级多线程又分为两种情况，一种是M:1的方式，M个用户线程对应一个内核进程。这种情况很容易因为一个系统阻塞，其他用户线程都会被阻塞，
无法利用机器多核的优势。还有一种模式就是M:N的方式，M个用户线程对应N个内核线程，这种模式一般需要语言运行时或库的支持，效率最高。

### 5.4.3 并发和调度
goroutine的好处：
+ (1) goroutine可以在用户空间调度，避免了内核态和用户态的切换导致的成本
+ (2) goroutine是语言原生支持的，提供了非常简洁的语法，屏蔽了大部分复杂底层实现
+ (3) goroutine更小的栈空间允许用户创建成千上万的实例。

Go的调度模型的三个实体：M、P、G

+ G是Go运行时对goroutine的抽象描述，G中存放并发执行的代码入口地址、上下文、运行环境（关联的P和M）、运行栈等执行相关的元信息。

+ M代表OS内核线程，是操作系统层面调度和执行的实体。M仅负责执行，M不停地被唤醒或创建，然后执行。M启动时进入的是运行时的管理代码
，由这段代码获取G和P资源，然后执行调度。另外，Go语言运行时会单独创建一个监控线程，负责对程序的内存、调度等信息进行监控和控制。

+ P代表M运行G所需要的资源，是对资源的一种抽象和管理，p不是一段代码实体，而是一个管理的数据结构，P主要是降低M管理调度G的复杂性，
增加一个间接的控制层数据库。P的数目默认是CPU核心的数量，G与P是M:N的关系，M可以成千上万，远大于N。


Go启动初始化过程
+ (1) 分配和检查栈空间
+ (2) 初始化参数和环境变量
+ (3) 当前运行线程标记为m0,m0是程序启动的主线程
+ (4) 调用运行时初始化函数runtime.schedinit进行初始化。主要是初始化内存空间分配器、GC、生成空闲P列表
+ (5) 在m0上调度第一个G,这个G运行runtime.main函数。runtime.main会拉起运行时的监控线程，然后调用main包的init()初始化函数，最后执行main函数。

抢占调度

`抢占调度的原因`
+ (1) 不让某个G长久地被系统调用阻塞，阻碍其他G运行。
+ (2) 不让某个G一直占用某个M不释放。
+ (3) 避免全局队列里面的G得不到执行。

`抢占调度的策略` 
+ (1)在进入系统调用(syscall)前后，各封装一层代码检测G的状态，当检测到当前G已经被监控线程抢占调度，则M停止执行当前G,进行调度切换。
+ (2)监控线程经过一段时间检测感知到P运行超过一定时间，取消P和M的关联，这也是一种更高层次的调度。 
+ (3)监控线程经过一段时间检测感知到G一直运行，超过了一定的时间，设置G标记，G执行栈扩展逻辑检测到抢占标记，根据相关条件决定是否抢占调度。




## 第六章 反射 day06
略
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
略