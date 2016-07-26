package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"InfiniteServer/models"
//	"fmt"
	"fmt"
	"strconv"
)

var dbhost = "43.241.236.248"
var dbport = "3306"
var dbuser = "skyun"
var dbpassword = "123123"
var dbname = "infinite"
//var dbconn = dbuser + ":" + dbpassword + "@tcp(" + dbhost + ")/" + dbname + "?charset=utf8"
var dbconn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbuser, dbpassword, dbhost, dbport, dbname)


type UserController struct {
	beego.Controller
}

func (this *UserController) init() {
	orm.RunSyncdb("default", false, true)
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

/**
 * 内部函数
 */
func auth(username string, password string) (*models.User, error) {
	user := new(models.User)
	err := initDB().QueryTable("user").Filter("username", username).Filter("password", password).One(user);
	return user, err
}

func getUser(uid string) (*models.User, error) {
	id, err := strconv.Atoi(uid)
	user := models.User{Id:id}
	o := initDB()
	err = o.Read(&user)
	o.LoadRelated(&user, "Objects");
	orm.RunSyncdb("default", false, true)
	return &user, err
}

func setAvatar(uid string, avatar string) error {
	id, err := strconv.Atoi(uid)
	user := models.User{Id:id}
	o := initDB();
	err = o.Read(&user)
	if (err == nil) {
		user.Avatar = avatar
		if num, err := o.Update(&user); err == nil {
			fmt.Println(num)
		}
	}
	return err
}

func initDB() orm.Ormer {
	orm.RegisterDriver("mysql", orm.DRMySQL);
	orm.RegisterDataBase("default", "mysql", dbconn);
	o := orm.NewOrm()
	o.Using("default")
	return o
}

func (this *UserController) getObjects() {
}

func TestGetUser() {
	user, err := auth("linyun", "aaa")
	if err != nil {
		print(err.Error());
	} else {
		print(user.Password)
	}
}

