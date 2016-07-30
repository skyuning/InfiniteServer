package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"InfiniteServer/models"
//	"fmt"
	"fmt"
	"strconv"
	"reflect"
)

var ormer orm.Ormer

type UserController struct {
	beego.Controller
}

func init() {
	ormer = models.GetOrmer()
	orm.RunSyncdb("default", false, true)
}

func auth(username string, password string) (*models.User, error) {
	user := new(models.User)
	err := ormer.QueryTable("user").Filter("username", username).Filter("password", password).One(user)
	return user, err
}

func getUser(uid string) (*models.User, error) {
	id, err := strconv.Atoi(uid)
	user := models.User{Id:id}
	err = ormer.Read(&user)
	ormer.QueryTable(new(models.Goods)).RelatedSel()
	ormer.LoadRelated(&user, "Objects", true)
	beego.Debug(reflect.TypeOf(user.Goods))
//	ormer.LoadRelated(&user.Objects[1], "Category");
	return &user, err
}

func setAvatar(uid string, avatar string) error {
	id, err := strconv.Atoi(uid)
	user := models.User{Id:id}
	err = ormer.Read(&user)
	if (err == nil) {
		user.Avatar = avatar
		if num, err := ormer.Update(&user); err == nil {
			fmt.Println(num)
		}
	}
	return err
}

/**
 * api
 */
func (this *UserController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")

	user, err := auth(username, password)
	if err == nil {
		this.Data["json"] = user
		this.ServeJSON()
	} else {
		this.Ctx.ResponseWriter.Status = 401
		this.Ctx.WriteString("auth failed: " + username)
	}
}

func (this *UserController) Get() {
	uid := this.Ctx.GetCookie("uid");
	user, err := getUser(uid)
	if (err == nil) {
		this.Data["json"] = user
		this.ServeJSON()
	} else {
		this.Ctx.ResponseWriter.Status = 401
		this.Ctx.WriteString(err.Error())
	}
}

func (this *UserController) SetAvatar() {
	uid := this.Ctx.GetCookie("uid");
	avatar := this.GetString("avatar")
	err := setAvatar(uid, avatar)
	resp := new(models.SimpleResponse)
	if (err == nil) {
		resp.Success = true
	} else {
		resp.Success = false
		resp.ErrMsg = err.Error()
	}

	this.Data["json"] = resp
	this.ServeJSON()
}

func TestGetUser() {
	user, err := auth("linyun", "aaa")
	if err != nil {
		print(err.Error());
	} else {
		print(user.Password)
	}
}

