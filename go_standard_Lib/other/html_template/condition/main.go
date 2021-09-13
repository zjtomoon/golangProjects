package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	//获取随机数，若不加随机种子，每次遍历获取都是重复的一些随机数据
	rand.Seed(time.Now().Unix()) //设置随机数种子，加上这行代码，可以保证每次随机都是随机的
	t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	sever := http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/process",process)
	sever.ListenAndServe()
}
