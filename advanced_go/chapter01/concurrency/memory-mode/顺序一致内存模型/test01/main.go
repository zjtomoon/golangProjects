package main

var (
	a string
	done bool
)

func setup()  {
	a = "hello,world"
	done = true
}

func main() {
	go setup()
	for !done {} //检测done变成true时，认为字符串初始化工作完成，然后进行字符串的打印工作
	print(a)
}
