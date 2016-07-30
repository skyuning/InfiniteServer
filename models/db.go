package models
import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

var (
	dbhost = "43.241.236.248"
	dbport = "3306"
	dbuser = "skyun"
	dbpassword = "123123"
	dbname = "infinite"
//	var dbconn = dbuser + ":" + dbpassword + "@tcp(" + dbhost + ")/" + dbname + "?charset=utf8"
	dbconn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbuser, dbpassword, dbhost, dbport, dbname)
)

func GetOrmer() orm.Ormer {
	orm.RegisterDriver("mysql", orm.DRMySQL);
	orm.RegisterDataBase("default", "mysql", dbconn);
	o := orm.NewOrm()
	o.Using("default")
	return o
}

