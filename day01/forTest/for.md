```go
//访问map
    for key ,value := range map{}
    for key := range map{}
    
//访问数组
    for index,value := range array{}
    for index := range array{}
    for _,value := range array{}
//访问切片
    for index,value := range slice{}
    for index := range slice{}
    for _,value := range slice{}
//访问通道
    for value := range channel{}
//类似C里面的for循环语句
for init;condition;post{}
//类似C里面的while循环语句
for condition{}
//类似C里面的while(1)死循环语句
for{}
```
