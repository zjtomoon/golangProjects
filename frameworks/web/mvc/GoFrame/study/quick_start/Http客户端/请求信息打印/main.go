package main

import (
	"bufio"
	"github.com/gogf/gf/frame/g"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		request := input.Text()
		response, err := g.Client().Get(request)
		if err != nil {
			panic(err)
		}
		response.RawDump()
	}

}
