package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:publishID`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["skitta.com/goweb/sinanewsAPI/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

}
