package funcs

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/net/gtcp"
	"golangProjects/frameworks/web/mvc/GoFrame/study/quick_start/网络服务开发/TCP组件/连接对象-通信开发进阶/types"
)

//自定义格式发送消息包
func SendPkg(conn *gtcp.Conn, act string, data ...string) error {
	s := ""
	if len(data) > 0 {
		s = data[0]
	}
	msg, err := json.Marshal(types.Msg{
		Act:  act,
		Data: s,
	})
	if err != nil {
		panic(err)
	}
	return conn.SendPkg(msg)
}

//自定义格式接收消息包
func RecvPkg(conn *gtcp.Conn) (msg *types.Msg, err error) {
	if data, err := conn.RecvPkg(); err != nil {
		return nil, err
	} else {
		msg = &types.Msg{}
		err = json.Unmarshal(data, msg)
		if err != nil {
			return nil, fmt.Errorf("invalid package structure: %s", err.Error())
		}
		return msg, err
	}
}
