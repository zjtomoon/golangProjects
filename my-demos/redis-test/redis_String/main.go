package main

// go get -u -v github.com/gomodule/redigo/redis
import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func main() {
	//通过 Go 向 Redis 写入数据和读取数据
	// 1. 连接到redis
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err = ",err)
		return
	}
	defer conn.Close()
	fmt.Println("connection succeeded !",conn)

	// 2. 通过 Go 向 Redis 写入数据 string[key-val]
	_,err = conn.Do("Set","name","贝尔·格里尔斯")
	if err != nil {
		fmt.Println("set err = ",err)
		return
	}
	fmt.Println("Successfully written !")

	// 3.通过 Go 向 Redis 读取数据 string[key-val]
	r, err := redis.String(conn.Do("Get","name"))
	if err != nil {
		fmt.Println("get err = ",err)
		return
	}
	// redis 包内部提供了转换函数，可以使用自带的转换函数。
	// 转换错误示例：
	//因为返回的 r 是 interface{},而 name 对应的值是 string,因此需要转换
	// nameString := r.(string)
	fmt.Printf("Read successfully!\n%s\n\n",r)
	multiSet()

}


//批量 Set/Get操作
func multiSet()  {
	// Connection
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err =",err)
		return
	}
	defer conn.Close()

	// Write
	_, err = conn.Do("MSET","name","Joe Chen","address","ShangHai")
	if err != nil {
		fmt.Println("MSET err = ",err)
		return
	}

	// Expire
	_, err = conn.Do("EXPIRE","address",10)
	if err != nil {
		fmt.Println("EXPIRE err = ",err)
		return
	}
	fmt.Println("Set automatic expiration!")


	// Rand
	// 必须使用 redis.Strings 函数将数组命令回复转换为字符串切片，否则无法遍历
	r, err := redis.Strings(conn.Do("MGET","name","address"))
	if err != nil {
		fmt.Println("MGET err = ",err)
	}

	for i, v := range r {
		fmt.Printf("%T\tr[%d] = %v\n",v,i,v)
	}

}
