package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "hello world !", Author: "Sau Sheong"},
		Post{Id: 2, Content: "Bonjour Monde", Author: "Pierre"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
	//由于程序在接下来的代码中就要立即对写入的posts.csv文件进行读取，而刚刚写入的数据有可能还滞留在缓冲区中
	//所以程序必须调用写入器的Flush方法来保证缓冲区中的所有数据都已经被正确地写入文件中

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	//将读取器的FieldsPerRecord字段设置为负数，这样即使读取器在读取时发现记录（record）里面缺少了某些字段，读取进程也不会被中断。
	//反之，如果FieldsPrRecord字段的值为正数，那么这个值就是用户要求从每条记录里面读取出的字段数量，当读取器从csv文件里面读取出的字段数量少于这个值时，就会抛出一个错误。
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
