# Kratos框架

## 官网

[A Go framework for microservices. | Kratos (go-kratos.dev)](https://go-kratos.dev/)

## 环境搭建

+ 环境依赖

位于阿里云盘/重装系统/sdks/go/kratos依赖

+ 安装

```bash
go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

+ 创建项目

```bash
# 创建项目模板
kratos new helloworld

cd helloworld
# 拉取项目依赖
go mod download
# 生成proto模板
kratos proto add api/helloworld/helloworld.proto
# 生成proto源码
kratos proto client api/helloworld/helloworld.proto
# 生成server模板
kratos proto server api/helloworld/helloworld.proto -t internal/service
```

+ 项目编译和运行

```bash
# 生成所有proto源码、wire等等
go generate ./...

# 运行项目
kratos run
```

+ 测试接口

```bash
curl 'http://127.0.0.1:8000/helloworld/kratos'

输出：
{
  "message": "Hello kratos"
}
```

