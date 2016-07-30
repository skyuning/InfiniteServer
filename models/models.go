package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(
		new(User),
		new(Goods),
		new(GoodsType),
	);
	ormer := GetOrmer()
	orm.RunSyncdb("default", false, true)

	// 衣
	ormer.Insert(&GoodsType{Name:"衣服"})
	// 食
	ormer.Insert(&GoodsType{Name:"食物"})
	// 住
	ormer.Insert(&GoodsType{Name:"帐篷"})
	// 行
	ormer.Insert(&GoodsType{Name:"野象"})

	orm.RunSyncdb("default", false, true)
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
	Goods    []*Goods `orm:"reverse(many)" json:"objects"`
}

type GoodsType struct {
	Id      int `orm:"pk;uniqu;auto" json:"id"`
	Name    string `orm:"unique" json:"name"`
	Objects []*Goods `orm:"reverse(many)" json:"objects,omitempty"`
}
type Goods struct {
	Id     int `orm:"pk;auto" json:"id"`
	Name   string `orm:"unique;null"`
	Amount int `json:"amount"`
	Type   *GoodsType `orm:"rel(fk)" json:"type"`
	User   *User `orm:"rel(fk);null;on_delete(do_nothing)" json:"-"`
}

