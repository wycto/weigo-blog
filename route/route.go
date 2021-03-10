package route

import (
	"github.com/wycto/weigo"
	"weigo-blog/app/controller/admin"
	"weigo-blog/app/controller/index"
	"weigo-blog/app/controller/public"
)

func init() {
	//路由定义必须遵循MVC：/public/user/  代表public应用（模块）、user控制器
	weigo.RouterStatic("/static/", "static")
	weigo.Router("/", &index.DefaultController{})
	weigo.Router("/public/user/", &public.UserController{})

	weigo.Router("/default/user/", &index.UserController{})
	weigo.Router("/default/api/", &index.APIController{})
	weigo.Router("/default/test/", &index.TestController{})
	weigo.Router("/default/model/", &index.ModelController{})

	weigo.Router("/admin/user/", &admin.UserController{})

}
