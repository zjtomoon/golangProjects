package main

import (
	"fmt"
	"sort"
	"time"
)

//学生成绩结构体
type StuScore struct {
	//姓名
	name string
	//成绩
	score int
}

type StuScores []StuScore

//Len()
func (s StuScores) Len() int {
	return len(s)
}

//Less():成绩将由低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

//Swap()
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}
	fmt.Println("Default :")
	//原始顺序
	for _, v := range stus {
		fmt.Println(v.name, ":", v.score)
	}

	startTime := time.Now()
	fmt.Println()
	//StuScores已经实现了sort.Interface接口
	sort.Sort(sort.Reverse(stus))
	fmt.Println("排序耗时：", time.Now().Sub(startTime))
	fmt.Println("Sorted :")
	//排好序的结构
	for _, v := range stus {
		fmt.Println(v.name, ":", v.score)
	}

	//判断是否已经排好顺序，将会打印true
	fmt.Println("Is Sorted?", sort.IsSorted(stus))

}
