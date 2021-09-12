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
	//遍历 allStudent
	for _,stu := range s.allStudent { //stu是具体的每一个学生
		fmt.Printf("学号：%d,姓名：%s\n",stu.id,stu.name)
	}
}

//增加学生
func (s studentMgr) addStudent() {
	//根据用户输入的内容创建一个新的学生
	var (
		StuId int64
		StuName string
	)
	//获取用户输入
	fmt.Println("请输入学号：")
	fmt.Scanln(&StuId)
	fmt.Println("请输入姓名")
	fmt.Scanln(&StuName)
	//根据用户输入创建结构体对象
	newStu := student{
		id: StuId,
		name: StuName,
	}
	//把新的学生放入到s.allStudent这个map中
	s.allStudent[newStu.id] = newStu
}

//修改学生
func (s studentMgr) editStudent() {
	//1、获取用户输入的学号
	var stuID int64
	fmt.Print("请输入学生id：")
	fmt.Scanln(&stuID)
	//2、展示该学号对应的学生信息,如果没有，提示”查无此人”
	stuObj,ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return 
	}
	//3、请输入修改后的学生名
	fmt.Printf("你要修改的学生信息如下:学号:%d 姓名:%s\n",stuObj.id,stuObj.name)
	fmt.Println("请输入学生的新名字：")
	var newName string
	fmt.Scanln(&newName)
	//4、更新学生的姓名
	stuObj.name = newName
	s.allStudent[stuID] = stuObj

}

//删除学生
func (s studentMgr) deleteStudent() {
	//1、请用户输入要删除的学生id
	var stuID int64
	fmt.Print("请输入要删除的学生的学号：")
	fmt.Scanln(&stuID)
	//2、去map找有没有这个学生，如果没有打印查无此人提示
	_,ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return 
	}

	//3、有的话就删除
	delete(s.allStudent,stuID)
	fmt.Println("删除成功！")
}
