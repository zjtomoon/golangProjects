package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

/*//变量验证
func main()  {
	validate := validator.New()
	//验证变量
	email := "admin#admin.com"
	//email := ""
	err := validate.Var(email,"required,email")
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		fmt.Println(validationErrors)
    	// output: Key: '' Error:Field validation for '' failed on the 'email' tag
    	// output: Key: '' Error:Field validation for '' failed on the 'required' tag
		return
	}
}
*/


type User struct {
	ID 		int64 	`json:"id" validate:"get=0"`
	Name 	string	`json:"name" validate:"required"`
	Gender	string	`json:"gender" validate:"required,oneof=man woman"`
	Age    	uint8  	`json:"age" validate:"required,gte=0,lte=130"`
	Email  	string 	`json:"email" validate:"required,email"`
}

//结构体验证
func main()  {
	validate := validator.New()
	user := &User{
		ID:1,
		Name:"frank",
		Gender:"boy",
		Age:18,
		Email:"gopher@88.com",
	}
	err := validate.Struct(user)
	if err != nil {
		validationErrrors := err.(validator.ValidationErrors)
		// output: Key: 'User.Age' Error:Field validation for 'Age' failed on the 'lte' tag
		// fmt.Println(validationErrors)
		fmt.Println(validationErrrors)
		return
	}
}