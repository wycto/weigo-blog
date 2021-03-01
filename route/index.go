package route

import (
	"github.com/wycto/weigo"
	"weigo-blog/app/controller/index"
)

func AddIndexRouter() {
	weigo.Router("/index/user/", &index.UserController{})
	weigo.Router("/index/api/", &index.APIController{})
	weigo.Router("/index/test/", &index.TestController{})
	weigo.Router("/index/model/", &index.ModelController{})
}
