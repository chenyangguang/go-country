package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Countries msg
var (
	Countries map[string]*Country
)

// Country  model
type Country struct {
	ID          int    `json:"id" orm:"column(id)"`
	Short2      string `json:"short2" orm:"column(short2);size(2)"`
	Short3      string `json:"short3" orm:"column(short3);size(3)"`
	LocalName   string `json:"local_name" orm:"column(local_name);size(255)"`
	EnglishName string `json:"english_name" orm:"column(english_name);size(100)"`
	PrefixPhone int    `json:"prefix_phone" orm:"column(prefix_phone)"`
}

func init() {
	orm.RegisterModel(new(Country))
}

// addOne Country
func addOne() {
	o := orm.NewOrm()
	country := Country{Short2: "CNH", Short3: "CNH", LocalName: "中国", EnglishName: "China", PrefixPhone: 86}
	id, err := o.Insert(&country)
	if err != nil {
		beego.Error(err)
	}

	beego.Error(id)
}
