# 安装go-zero库
```bash

GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/tal-tech/go-zero

```


# 安装goctl工具
```bash

 # any version 
 GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/tal-tech/go-zero/tools/goctl
   
 # go1.16 and newer   
 GOPROXY=https://goproxy.cn/,direct go install github.com/tal-tech/go-zero/tools/goctl@latest

```

# 快速开始

```bash
     goctl api new greet
     cd greet
     go mod init
     go mod tidy
     go run greet.go -f etc/greet-api.yaml
```

# 官方文档

[go-zero官方文档](https://go-zero.dev/cn/)