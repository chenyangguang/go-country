package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/chenyangguang/go-country/controllers:CountryController"] = append(beego.GlobalControllerRouter["github.com/chenyangguang/go-country/controllers:CountryController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/chenyangguang/go-country/controllers:CountryController"] = append(beego.GlobalControllerRouter["github.com/chenyangguang/go-country/controllers:CountryController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/get_all`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/chenyangguang/go-country/controllers:CountryController"] = append(beego.GlobalControllerRouter["github.com/chenyangguang/go-country/controllers:CountryController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/chenyangguang/go-country/controllers:CountryController"] = append(beego.GlobalControllerRouter["github.com/chenyangguang/go-country/controllers:CountryController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/chenyangguang/go-country/controllers:CountryController"] = append(beego.GlobalControllerRouter["github.com/chenyangguang/go-country/controllers:CountryController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

}
