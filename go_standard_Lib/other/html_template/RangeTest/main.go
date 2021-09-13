package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter,r *http.Request)  {
	t,_ := template.ParseFiles("index.html")
	dayOfWeek := []string{"Mon","Tue","Wed","Thu","Fri","Sat","Sun"}
	t.Execute(w,dayOfWeek)
}

func main() {
	sever := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	sever.ListenAndServe()
}
