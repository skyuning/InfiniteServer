package models

import (
	"github.com/astaxie/beego/orm"
)


type User struct {
	Id       int `orm:"pk;auto" json:"id"`
	Username string `json:"username"`
	Password string `json:"p,omitempty"`
}

func (this *User) GetPassword() string {
	return this.Password;
}

func init() {
	orm.RegisterModel(new(User));
}
