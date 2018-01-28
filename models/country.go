package models

import (
	"github.com/astaxie/beego/orm"
)

// Countries msg
var (
	Countries map[int]*Country
)

// Country  model
type Country struct {
	ID          int64  `json:"id" orm:"pk;column(id)"`
	Short2      string `json:"short2" orm:"column(short2);size(2)"`
	Short3      string `json:"short3" orm:"column(short3);size(3)"`
	LocalName   string `json:"local_name" orm:"column(local_name);size(255)"`
	EnglishName string `json:"english_name" orm:"column(english_name);size(100)"`
	PrefixPhone int64  `json:"prefix_phone" orm:"column(prefix_phone)"`
}

// TableName setting
func (c *Country) TableName() string {
	return "country"
}

// TableEngine ...
func (c *Country) TableEngine() string {
	return "INNODB"
}

// TableIndex setting
func (c *Country) TableIndex() [][]string {
	return [][]string{
		[]string{"LocalName", "EnglishName"},
	}
}

// TableUnique setting
func (c *Country) TableUnique() [][]string {
	return [][]string{
		[]string{"Short2", "Short3", "LocalName", "EnglishName", "PrefixPhone"},
	}
}

func init() {
	orm.RegisterModel(new(Country))
}

// GetAllCountries get all the countries info.
func GetAllCountries() []*Country {
	o := orm.NewOrm()
	var countries []*Country
	q := o.QueryTable("country")
	_, err := q.All(&countries)
	if err != nil {
		return nil
	}
	return countries
}

// GetCountryByID ...
func GetCountryByID(id int64) (Country, bool) {
	c := Country{ID: id}
	o := orm.NewOrm()
	err := o.Read(&c)
	if err != nil {
		return c, false
	}
	return c, true
}
