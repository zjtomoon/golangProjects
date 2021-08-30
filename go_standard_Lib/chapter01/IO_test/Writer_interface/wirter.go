package Writer_interface

import (
	"fmt"
	"net/http"
)

func Index(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "hello world")
}
