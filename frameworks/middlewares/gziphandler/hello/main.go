package main

import (
	"github.com/NYTimes/gziphandler"
	"io"
	"net/http"
)

func main() {
	withoutGZ := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Conten-Type", "text/plain")
		io.WriteString(w, "Hello,World")
	})
	withGZ := gziphandler.GzipHandler(withoutGZ)

	http.Handle("/", withGZ)
	http.ListenAndServe(":8080", nil)
}
