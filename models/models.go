package models

import (
	"github.com/astaxie/beego/orm"
)


type User struct {
	Id       int `orm:"pk;auto"`
	Username string
	Password string
}

func init() {
	orm.RegisterModel(new(User));
}
