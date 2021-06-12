package main
//go get -u -v github.com/gorilla/websocket
//go get -u -v github.com/gorilla/mux
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws",myws)
	if err := http.ListenAndServe("127.0.0.1:8181",router); err != nil{
		fmt.Println("err:",err)
	}
}