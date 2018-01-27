package controllers

import (
	"github.com/astaxie/beego"
)

// CountryController operation about country
type CountryController struct {
	beego.Controller
}

// Post creat a country
func (o *CountryController) Post() {
	o.ServeJSON()
}

// Get test
func (o *CountryController) Get() {
	o.ServeJSON()
}

// GetAll countries
func (o *CountryController) GetAll() {
	o.ServeJSON()
}

// Delete a country
func (o *CountryController) Delete() {
	o.ServeJSON()
}
