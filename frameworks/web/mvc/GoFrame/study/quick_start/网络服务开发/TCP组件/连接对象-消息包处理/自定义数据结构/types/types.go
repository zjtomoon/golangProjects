package types

import "github.com/gogf/gf/frame/g"

type NodeInfo struct {
	Cpu      float32
	Host     string
	Ip       g.Map
	MemUsed  int
	MemTotal int
	Time     int
}
