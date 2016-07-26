package routers

import (
	"InfiniteServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/login", &controllers.UserController{}, "*:Login")
	beego.Router("/user/set_avatar", &controllers.UserController{}, "post:SetAvatar")
}
