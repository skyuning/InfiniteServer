package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"InfiniteServer/models"
//	"fmt"
	"fmt"
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

/**
 * api
 */
func (this *UserController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")

	user, err := auth(username, password)
	if err != nil {
		this.Ctx.ResponseWriter.Status = 401
		this.Ctx.WriteString("auth failed" + username)
		return
	}

	this.Data["json"] = user;
	this.ServeJSON();
}

/**
 * 内部函数
 */
func auth(username string, password string) (*models.User, error) {
	user := new(models.User)
	err := initDB().QueryTable("user").Filter("username", username).Filter("password", password).One(user);
	user.Password = ""
	return user, err
}

func initDB() orm.Ormer {
	orm.RegisterDriver("mysql", orm.DRMySQL);
	orm.RegisterDataBase("default", "mysql", dbconn);
	o := orm.NewOrm()
	o.Using("default")
	return o
}

func TestGetUser() {
	user, err := auth("linyun", "aaa")
	if err != nil {
		print(err.Error());
	} else {
		print(user.GetPassword())
	}
}

