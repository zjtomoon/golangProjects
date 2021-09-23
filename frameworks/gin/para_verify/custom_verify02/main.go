package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"time"
)

type Booking struct {
	//定义一个预约的时间大于今天的时间
	CheckIn		time.Time	`form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut	time.Time	`form:"check_out" binding:"required,gtfield=checkIn" time_format:"2006-01-02"`
}

func bookableDate(
	v *validator.Validate,topStruct reflect.Value,currentStructOrField reflect.Value,
	field reflect.Value,fieldType reflect.Type,fieldKind reflect.Kind,param string,
) bool {
	//field.Interface().(time.Time)获取参数值并且转换成时间格式
	if data,ok := field.Interface().(time.Time);ok {
		today := time.Now()
		if today.Unix() > data.Unix() {
			return false
		}
	}
	return true
}

func getBookable(c *gin.Context)  {
	var b Booking
	if err := c.ShouldBindWith(&b,binding.Query);err == nil{
		c.JSON(http.StatusOK,gin.H{"message":"Booking dates are vaild"})
	} else {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
}

func main()  {
	route := gin.Default()
	//注册验证
	if v,ok := binding.Validator.Engine().(*validator.Validate);ok {
		//绑定第一个参数是验证的函数，第二个参数是自定义的验证函数
		v.RegisterValidation("bookabledate",bookableDate)
	}
	route.GET("/51mh",getBookable)
	route.Run()
}
// curl -X GET "http://localhost:8080/5lmh?check_in=2019-11-07&check_out=2019-11-20"
// curl -X GET "http://localhost:8080/5lmh?check_in=2019-09-07&check_out=2019-11-20"
// curl -X GET "http://localhost:8080/5lmh?check_in=2019-11-07&check_out=2019-11-01"