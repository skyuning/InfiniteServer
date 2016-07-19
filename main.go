package main

import (
	_ "InfiniteServer/routers"
	"github.com/astaxie/beego"
)

func main() {
	//	controllers.TestGetUser()
	beego.Run()
}

