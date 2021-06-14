package main

import "fmt"

//有一个物件:
//1、保存了一些数据 --->结构体的字段
//2、有三个功能     --->结构体的方法
type student struct {
	id   int64
	name string
}

//造一个学生的管理者
type studentMgr struct {
	allStudent map[int64]student
}

//查看学生
func (s studentMgr) showStudents() {
	for _,stu := range s.allStudent {
		fmt.Printf("学号：%d,姓名：%s\n",stu.id,stu.name)
	}
}

//增加学生
func (s studentMgr) addStudent() {
	
}

//修改学生
func (s studentMgr) editStudent() {

}

//删除学生
func (s studentMgr) deleteStudent() {

}
