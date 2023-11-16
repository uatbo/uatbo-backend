package main

import (
	"fmt"
	"uatbo-backend/core"
	"uatbo-backend/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 连接数据库
	global.DB = core.InitGorm()
	fmt.Println(global.Config)

}
