package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"bluebell/settings"
	"fmt"
)

func main() {

	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := snowflake.Init(settings.Conf.StartTime, uint16(settings.Conf.MachineID)); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}

	if err := mysql.Init(settings.Conf); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	r := router.SetupRouter()
	r.Run(fmt.Sprintf(":%d", settings.Conf.Port))

}
