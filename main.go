package main

import "bluebell/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
