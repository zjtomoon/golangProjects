package main

import (
	"fmt"
)

//定义一个HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode //这个表示指向前一个节点
	next     *HeroNode //这个表示指向下一个节点
}

//给双向链表插入一个节点
//编写第一种插入方法，在最后加入【简单】
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//思路
	//1. 先找到该链表地最后这个节点
	//2. 创建一个辅助节点
	temp := head
	for {
		if temp.next == nil { //表示找到最后
			break
		}
		temp = temp.next //让temp不断的指向下一个节点
	}

	//3. 将newHeroNode加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

//给双向链表插入一个节点
//编写第2种插入方法，根据no 的编号从小到大插入..【实用】
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//思路
	//1. 先找到该链表的最后这个节点
	//2. 创建一个辅助节点
	temp := head
	flag := true
	for {
		if temp.next == nil { //说明到链表的最后
			break
		} else if temp.next.no >= newHeroNode.no {
			//说明newHeroNode 就应该插入到temp后面
			break
		} else if temp.next.no == newHeroNode.no {
			//说明链表中已经有了这个no,就不让插入
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("对不起，已经存在no=", newHeroNode)
		return
	} else {
		newHeroNode.next = temp.next //ok
		newHeroNode.pre = temp       //ok
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}
}

//删除一个节点【双向链表删除一个节点】
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	//找到要删除节点的no,和temp的下一个节点的no比较
	for {
		if temp.next == nil { //说明到链表的最后
			break
		} else if temp.next.no == id {
			//说明找到了
			flag = true
			break
		}
		temp = temp.next
	}
	if flag { //找到，删除
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("抱歉，要删除的id不存在")
	}
}

//显示链表的所有节点信息
//这里仍然使用单向的链表显示方式
func ListHeroNode(head *HeroNode) {
	//1. 创建一个辅助节点
	temp := head

	//先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也...")
		return
	}
	//2. 遍历这个链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		//判断是否链表后
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func ListHeroNode2(head *HeroNode) {
	//1. 创建一个辅助节点
	temp := head

	//先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也...")
		return
	}
	//2. 让temp定位到双向链表的最后节点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	//3. 遍历这个链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.no, temp.name, temp.nickname)
		//判断是否链表头
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

func main() {
	//1. 先创建一个头节点
	head := &HeroNode{}

	//2. 创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero2)
	ListHeroNode(head)
	fmt.Println()
	fmt.Println("逆序打印")
	ListHeroNode2(head)

}
