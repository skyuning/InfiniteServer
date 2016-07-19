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

func (this *UserController) Get() {
	username := this.GetString("username")

	user, err := GetUser(username)
	if err != nil {
		this.Ctx.WriteString("no user named " + username)
		return
	}

	this.Data["json"] = user
	this.ServeJSON()
}

func (this *UserController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	user, err := GetUser(username)
	if err != nil {
		this.Ctx.WriteString("no user named " + username)
		return
	}

	if password == user.Password {
		this.Ctx.WriteString(fmt.Sprintf("Login Success: %s", username))
	} else {
		this.Ctx.WriteString(fmt.Sprintf("Incorrect Password: %s", password))
	}
}

func GetUser(username string) (*models.User, error) {
	user := new(models.User)
	err := initDB().QueryTable("user").Filter("username", username).One(user);
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
	user, err := GetUser("linyun")
	if err != nil {
		print(err.Error());
	} else {
		print(user.Password)
	}
}

