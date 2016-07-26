package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User), new(Object));
}

type SimpleResponse struct {
	Success bool `json:"success"`
	ErrMsg  string `json:"err_msg"`
}

type User struct {
	Id       int `orm:"pk;auto" json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Avatar   string `json:"avatar"`
	Objects []*Object `orm:"reverse(many)" json:"object,omitempty"`
}

type Object struct {
	Id   int `orm:"pk;auto" json:"id"`
	Name string `json:"name"`
	Amount int `json:"amount"`
	User *User `orm:"rel(fk);null;on_delete(do_nothing)" json:"-"`
}

