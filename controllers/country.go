package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/chenyangguang/go-country/models"
)

// CountryController operation about country
type CountryController struct {
	beego.Controller
}

// Post creat a country
func (c *CountryController) Post() {
	var co models.Country

	short2 := c.GetString("short2")
	short3 := c.GetString("short3")
	cn := c.GetString("local_name")
	en := c.GetString("english_name")
	ph, _ := c.GetInt64("prefix_phone", 0)

	if cn == "" {
		c.Ctx.WriteString("localname is empty")
		return
	}
	if en == "" {
		c.Ctx.WriteString("english is empty")
		return
	}
	if short2 == "" {
		c.Ctx.WriteString("short2 is empty")
		return
	}
	if ph == 0 {
		c.Ctx.WriteString("prefix phone is empty")
		return
	}
	co.Short2 = short2
	co.Short3 = short3
	co.LocalName = cn
	co.EnglishName = en
	co.PrefixPhone = ph

	o := orm.NewOrm()
	existed := o.QueryTable("country").Filter("Short2", short2).Filter("EnglishName", en).Exist()
	if existed {
		c.Data["json"] = "Exsisted Country"
	} else {
		cid := models.AddCountry(&co)
		if cid == 0 {
			c.Data["json"] = 0
		}
		c.Data["json"] = cid
	}
	c.ServeJSON()
}

// Get test
func (c *CountryController) Get() {
	id := c.Ctx.Input.Param(":id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	if intID != 0 {
		s, err := models.GetCountryByID(intID)
		if !err {
			c.Data["json"] = "Not found"
		} else {
			c.Data["json"] = s
		}
	}
	c.ServeJSON()
}

// GetAll countries
func (c *CountryController) GetAll() {
	cs := models.GetAllCountries()
	if cs == nil {
		c.Data["json"] = "Not found."
	} else {
		c.Data["json"] = cs
	}
	c.ServeJSON()
}

// Delete a country
func (c *CountryController) Delete() {
	id := c.Ctx.Input.Param(":id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	o := orm.NewOrm()
	if _, err := o.QueryTable("country").Filter("ID", intID).Delete(); err != nil {
		c.Data["json"] = map[string]interface{}{"result": false, "msg": "Operation fail!"}
	} else {
		c.Data["json"] = map[string]interface{}{"result": true, "msg": "Operation success"}
	}
	c.ServeJSON()
}

// Put for update country
func (c *CountryController) Put() {
	id := c.Ctx.Input.Param(":id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	co := models.Country{ID: intID}
	o := orm.NewOrm()
	if o.Read(&co) == orm.ErrNoRows {
		c.Data["json"] = map[string]interface{}{"result": false, "msg": "Not Found!"}
		c.ServeJSON()
	}

	co.Short2 = c.GetString("short2")
	co.Short3 = c.GetString("short3")
	co.LocalName = c.GetString("local_name")
	co.EnglishName = c.GetString("english_name")
	ph, _ := c.GetInt64("prefix_phone", 0)
	co.PrefixPhone = ph

	res, err := models.UpdateCountry(intID, &co)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = res
	}
	c.ServeJSON()
}
