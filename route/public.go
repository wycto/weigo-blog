package route

import (
	"github.com/wycto/weigo"
	"weigo-blog/app/controller/public"
)

func AddPublicRouter() {
	weigo.Router("/public/user/", &public.UserController{})
}
