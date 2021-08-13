package main

import "fmt"


//约瑟夫问题

type Boy struct {
	No   int
	Next *Boy //指向下一个小孩的指针（默认值是nil)
}

// 编写一个函数，构成单向的环形链表
// num : 表示小孩的个数
// *Boy : 返回该环形的链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}  //空节点
	curBoy := &Boy{} //空节点

	//判断
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}

	//循环的构建这个环形链表
	for i := 0; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		//分析构成循环链表，需要一个辅助指针
		//1. 因为第一个比较特殊
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first //构造环形链表
		}
	}

	return first
}
//显示单向的环形链表[遍历]
func ShowBoy(first *Boy)  {

	//处理一下如果环形链表为空
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩...")
		return
	}

	//创建一个指针，帮助遍历
	curBoy := first
	for {
		fmt.Printf("小孩编号= %d -> ",curBoy.No)
		//退出的条件?curBoy.Next == first
		if curBoy.Next == first {
			break
		}
	}
}

func main() {

}