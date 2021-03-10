package admin

import "github.com/wycto/weigo"

type UserController struct {
	weigo.Controller
}

func (c *UserController) Login() {
	c.Display("")
}
