# gin实战-基于gin实现动态博客

[视频地址](https://www.bilibili.com/video/av73698322?t=2400&p=5)

## 项目结构 

+ model：实体 

+ database：数据层 

+ service：业务逻辑 

+ controller：页面控制相关 

+ static：css js 
+ static/html：html模板

+ utils：工具 

+ main.go：入口，定义路由

## 数据库分析

+ article ：文章表
+ category：分类表
+ comment：评论表
+ leave：留言表

## 主页的实现思路

### 1、实体类

+ 分类的结构体
+ 文章的结构体

### 2、数据层

+ init()数据库初始化函数
+ 分类相关的操作（添加，查询，查1个分类，查多个分类，查所有分类）
+ 文章相关的操作（添加文章（投稿），查询所有文章，根据文章id查看内容）

