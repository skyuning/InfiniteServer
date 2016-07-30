package main

import (
	_ "InfiniteServer/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogger("console", "")
	//	controllers.TestGetUser()
	beego.Run()
}

