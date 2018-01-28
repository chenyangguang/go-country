package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/chenyangguang/go-country/models"
)

// CountryController operation about country
type CountryController struct {
	beego.Controller
}

// Post creat a country
func (c *CountryController) Post() {
	c.ServeJSON()
}

// Get test
func (c *CountryController) Get() {
	id := c.Ctx.Input.Param(":id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	s, err := models.GetCountryByID(intID)
	if !err {
		c.Data["json"] = "Not found"
	} else {
		c.Data["json"] = s
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
	c.ServeJSON()
}
