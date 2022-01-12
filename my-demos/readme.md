# 笔记

+ go tool编译和查看汇编代码

```bash
# 生成main.o
go tool compile .\main.go
# 查看汇编
go tool objdump .\main.o
```

