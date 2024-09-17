package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/router"
	"bluebell/settings"
	"fmt"
)

func main() {

	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
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
	r.Run(":80")

}
