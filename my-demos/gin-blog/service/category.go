package service

import (
	"golangProjects/my-demos/gin-blog/database/db"
	"golangProjects/my-demos/gin-blog/model"
)

//获取所有分类
func GetALLCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		return
	}
	return
}
