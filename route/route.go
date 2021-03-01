package route

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wycto/weigo"
	"weigo-blog/app/controller/index"
	"weigo-blog/app/controller/public"
)

func init() {
	//路由定义必须遵循MVC：/public/user/  代表public应用（模块）、user控制器
	weigo.Router("/", &index.DefaultController{})
	weigo.Router("/public/user/", &public.UserController{})

	weigo.Router("/index/user/", &index.UserController{})
	weigo.Router("/index/api/", &index.APIController{})
	weigo.Router("/index/test/", &index.TestController{})
	weigo.Router("/index/model/", &index.ModelController{})

	httprouter.New().POST("/public/test/:action", nil)

}
