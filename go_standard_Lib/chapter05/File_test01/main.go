package main

import (
	"os"
	"log"
)

func main() {
	file, err := os.Create("./studygolang.txt")
	if err != nil {
		log.Fatal("error:",err)
	}
	defer file.Close()

	fileMode := getFileMode(file)
	log.Println("file mode :",fileMode)
	file.Chmod(fileMode | os.ModeSticky)

	log.Println("change after,file mode :",getFileMode(file))
}

func getFileMode(file *os.File) os.FileMode {
	fileInfo,err := file.Stat()
	if err != nil {
		log.Fatal("file stat error :",err)
	}
	return fileInfo.Mode()
}

