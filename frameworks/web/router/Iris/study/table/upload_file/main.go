package main

import (
	"github.com/kataras/iris/v12"
	"mime/multipart"
	"strings"
)

func main() {
	app := iris.Default()
	app.Post("/upload", iris.LimitRequestBodySize(1000), func(ctx iris.Context) {
		//
		// UploadFormFiles
		// uploads any number of incoming files ("multiple" property on the form input).
		//

		// The second, optional, argument
		// can be used to change a file's name based on the request,
		// at this example we will showcase how to use it
		// by prefixing the uploaded file with the current user's ip.
		//ctx.UploadFormFiles("./uploads", beforeSave)
	})

	app.Run(iris.Addr(":8080"))
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	// make sure you format the ip in a way
	// that can be used for a file name (simple case):
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)

	// you can use the time.Now, to prefix or suffix the files
	// based on the current time as well, as an exercise.
	// i.e unixTime :=    time.Now().Unix()
	// prefix the Filename with the $IP-
	// no need for more actions, internal uploader will use this
	// name to save the file into the "./uploads" folder.
	file.Filename = ip + "-" + file.Filename
}
