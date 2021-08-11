package main

import "fmt"

//定义一个HeroNode
type HeroNode struct {
	no			int
	name		string
	nickname	string
	next		*HeroNode //这个表示指向下一个节点
}

//给链表插入一个节点
//编写第一种插入方法，在单链表的最后加入【简单】
func InsertHeroNode(head *HeroNode,newHeroNode *HeroNode)  {
	//思路
	//1. 先找到该链表的最后这个节点
	//2. 创建一个辅助节点
	temp := head
	for  {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	//3. 将newHeroNode加入到链表的最后
	temp.next = newHeroNode
}

//给链表插入一个节点
//编写第二种插入方法，根据no的编号从小到大插入 【实用】
func InsertHeroNode2(head *HeroNode,newHeroNode *HeroNode)  {
	//思路
	//1. 找到适当的节点
	//2. 创建一个辅助节点
	temp := head
	flag := true
	//让插入的节点的no,和temp的下一个节点的no比较
	for {
		if temp.next == nil {
			break
		} else if temp.next.no >= newHeroNode.no {
			//说明newHeroNode就应该插入到temp后面
			break
		}else if temp.next.no == newHeroNode.no {
			//说明我们的链表中已经有这个no,就不用插入了
			flag = false
			break
		}

		temp = temp.next
	}

	if !flag {
		fmt.Println("对不起，已经存在no= ",newHeroNode.no)
		return
	}else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

func main() {

}