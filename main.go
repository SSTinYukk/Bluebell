package main

import (
	"bluebell/controller"
	"bluebell/router"
	"fmt"
)

func main() {

	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("inti validator trans failed, err:%v\n", err)
		return
	}

	r := router.SetupRouter()
	r.Run(":80")

}
