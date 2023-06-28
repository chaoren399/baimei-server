package main

import (
	"baimei-server/global"
	"baimei-server/initiallize"
)

// 项目启动入口
func main() {
	r := initiallize.Routers()
	initiallize.Viper()
	initiallize.K8SWithDiscovery()
	panic(r.Run(global.CONF.System.Addr))
}
