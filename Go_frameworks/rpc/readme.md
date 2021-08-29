# rpc框架

## 安装

```bash
go get -u -v github.com/smallnest/rpcx/...
```

这一步只会安装 rpcx 的基础功能。如果你想要使用 etcd 作为注册中心，你需要加上`etcd`这个标签。 (see [build tags](https://golang.org/pkg/go/build/#hdr-Build_Constraints))

```go
go get -u -v -tags "etcd" github.com/smallnest/rpcx/...
```

如果你想要使用 quic ，你也需要加上`quic`这个标签。

```go
go get -u -v -tags "quic etcd" github.com/smallnest/rpcx/...
```

方便起见，我推荐你安装所有的tags，即使你现在并不需要他们：

```go
go get -u -v -tags "reuseport quic kcp zookeeper etcd consul ping" github.com/smallnest/rpcx/...
```

