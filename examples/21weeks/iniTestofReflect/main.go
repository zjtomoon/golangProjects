package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// ini配置文件解析器

// MysqlConfig mysql配置结构体
type MysqlConfig struct {
	Address 	string	`ini:"address"`
	Port 		int		`ini:"port"`
	Username 	string	`ini:"username"`
	Password	string	`ini:"password"`
}

//RedisConfig redis配置结构体
type RedisConfig struct {
	Host 		string	`ini:"host"`
	Port		int		`ini:"port"`
	Password	string	`ini:"password"`
	Database	int		`ini:"database"`
}
// Config ...
type Config struct {
	MysqlConfig	`ini:"mysql"`
	RedisConfig	`ini:"redis"`
}

func loadIni(fileName string,data interface{}) (err error) {
	 // 0. 参数地校验
	 // 0.1 传进来地data参数必须是指针类型(因为需要在函数中对其赋值)
	 t := reflect.TypeOf(data)
	 fmt.Println(t,t.Kind())
	 if t.Kind() != reflect.Ptr {
		 //err = fmt.Errorf("data should be a pointer") //格式化输出之后返回一个error类型
		 err = errors.New("data param should be a pointer") //新建一个错误
		 return
	 }
	 //	0.2 传进来地data参数必须是结构体指针(因为配置文件中各种键值对需要赋值给结构体字段)
	 if t.Elem().Kind() != reflect.Struct {
		 err = errors.New("data param should be a struct pointer") //新建一个错误
		 return
	 }
	 // 1. 读文件得到字节类型数据
	 b,err := ioutil.ReadFile(fileName)
	 if err != nil {
		 return
	 }
	 //string(b) //将字节类型地文件内容转换成字符串
	 lineSlice := strings.Split(string(b),"\n")
	 fmt.Printf("%#v\n",lineSlice)
	 // 2. 一行一行地读数据
	 var structName string
	 for idx,line := range lineSlice {
		 // 去掉字符串首尾地空格
		 line = strings.TrimSpace(line)
		 // 2.1 如果是注释就跳过
		 if strings.HasPrefix(line,";") || strings.HasPrefix(line,"#") {
			 continue
		 }
	 	// 2.2 如果是[开头地就表示是节(section)
		 if strings.HasPrefix(line,"[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error",idx + 1)
				return
			}
			//把这一行首尾地[]去掉，取到中间的内容把首尾的空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1:len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error",idx + 1)
				return
			}
			//根据字符串sectionName去data里面根据反射找到对应的结构体
			//v := reflect.ValueOf(data)
			for i := 0 ; i < t.Elem().NumField() ; i++ {
				field := t.Field(i)
				if sectionName == field.Tag.Get("ini") {
					//说明找到了对应的嵌套结构体，把字段名记下来
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n",sectionName,structName)
					break
				}
			}
		 } else {
			// 2.3 如果不是[开头就是分割地键值对
			// 1、以等号分割这一行，等号左边是key,等号右边是value
			
			// 2、根据structName 去 data里面把对应的嵌套结构体取出来
			// 3、遍历嵌套结构体的每一个字段，判断tag是不是key
			// 4、如果key == tag,给这个字段赋值

		 }

	 	
	 }
	 return
}

func main()  {
	var cfg Config
	err := loadIni("./config.ini",&cfg)
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n",err)
		return
	}
	fmt.Println(cfg)
}