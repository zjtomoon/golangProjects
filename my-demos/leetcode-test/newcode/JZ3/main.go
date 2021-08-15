/*package main

import "sort"

  type ListNode struct{
    Val int
    Next *ListNode
  }


func printListFromTailToHead( head *ListNode ) []int {
	// write code here
	temp := head
	res := []int{}
	//var answer []int
	for {
		res = append(res, temp.Val)
		temp = temp.Next
		if temp.Next == nil {
			res = append(res, temp.Val)
			break
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	return res
}*/
package main
import . "nc_tools"
import "container/list"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param head ListNode类
 * @return int整型一维数组
*/
func printListFromTailToHead( head *ListNode ) []int {
	// write code here
	l := list.New()
	length := 0
	for ; head != nil ;head = head.Next {
		l.PushFront(head.Val)
		length++
	}

	answer := make([]int,length)
	item := l.Front()
	for i:=0;i<length;i++ {
		answer[i]=(item.Value).(int)
		item = item.Next()
	}

	return answer

}
/*
package main
import . "nc_tools"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param head ListNode类
 * @return int整型一维数组
*/
func printListFromTailToHead2( head *ListNode ) []int {
	// write code here
	var answer []int

	if head == nil {
		return answer
	}
	temp := printListFromTailToHead(head.Next)
	answer = append(answer,temp...)
	answer = append(answer,head.Val)
	return answer
}