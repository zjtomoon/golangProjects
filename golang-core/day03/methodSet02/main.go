package main

type Data struct{}

func (Data) TestValue()  {}

func (*Data) TestPointer() {}
//这种字面量显式调用，无论值调用还是表达式调用，编译器都不会进行方法集的自动转换，编译器会严格校验方案

//*Data方法集是TestPointer 和 TestValue
//Data方法集只有TestValue

(*Data) (&struct{}{}).TestPointer() //显式的调用
(*Data) (&struct{}{}).TestValue()	//显式的调用

(Data) (struct{}{}).TestValue() //method value
Data.TestValue(struct{}{})

//如下调用因为方法集和不匹配而失败
//Data.TestPointer(struct{}{})
//(Data) (struct{}{}).TestPointer()
Data(struct {} literal)

var a Data = struct{}{}

//表达式调用编译器不会进行自动转换

Data.TestValue(a)
//Data.TestValue(&a)
(*Data).TestPointer(&a)
Data.TestPointer(&a)

//值调用编译器会进行自动转换
f := a.TestValue
f()

y := (&a).TestValue //编译器帮助转换a.TestValue
y()

g := a.TestPointer //会转换为(&a).TestPointer
g()

x := (&a).TestPointer
x()

