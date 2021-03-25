package main

import (
	"io"
	"os"
)

func CopyFile(dst, src string) (w int64, err error) {
	src, err := os.Open(src)
	if err != nil {
		return
	}
	dst, err := os.Create(dst)
	if err != nil {
		//src容易被忘记关闭
		src.Close()
		return
	}
	w, src = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}
