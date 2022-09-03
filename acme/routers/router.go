package routers

import (
	"acme/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/home", &controllers.MainController{})
	beego.Router("/user/login/:back", &controllers.MainController{}, "get,post:Login")
	beego.Router("/user/logout", &controllers.MainController{}, "get:Logout")
	beego.Router("/user/register", &controllers.MainController{}, "get,post:Register")
	beego.Router("/user/profile", &controllers.MainController{}, "get,post:Profile")
	beego.Router("/user/verify/:uuid", &controllers.MainController{}, "get:Verify")
	beego.Router("/user/remove", &controllers.MainController{}, "get,post:Remove")
	beego.Router("/notice", &controllers.MainController{}, "get:Notice")
}
