package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := int32(97)
	str := "10000"
	ret1,err := strconv.ParseInt(str,10,64)
	if err != nil {
		fmt.Println("parseint failed,err:",err)
		return
	}
	fmt.Printf("%#v %T\n",ret1,ret1)
	fmt.Printf("%#v %T\n",ret1,int(ret1))
	retInt,_ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n",retInt,retInt)

	ret2 := fmt.Sprintf("%d",i) //"97"
	fmt.Printf("%#v",ret2)

	ret3 := strconv.Itoa(int(i))
	fmt.Printf("%#v %T\n",ret3,ret3)
	
	//从字符串中解析出布尔值
	boolStr := "true"
	boolValue,_ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n",boolValue,boolValue)

	//从字符串解析出浮点数
	floatStr := "1.234"
	floatValue,_ := strconv.ParseFloat(floatStr,64)
	fmt.Printf("%#v %T\n",floatValue,floatValue)

}
