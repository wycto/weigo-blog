package index

import (
	"github.com/wycto/weigo"
	"io"
	"weigo-blog/app/model"
)

type UserController struct {
	weigo.Controller
	modelName string
}

func (c UserController) Index() {
	io.WriteString(c.Context.ResponseWriter, "{aa:12}")
}

func (c *UserController) Login() {
	io.WriteString(c.Context.ResponseWriter, "this is user login"+c.Context.Request.Method)
	io.WriteString(c.Context.ResponseWriter, "name:"+c.Context.Request.FormValue("name"))
	io.WriteString(c.Context.ResponseWriter, "name:"+c.Context.Request.Form.Get("name"))
	io.WriteString(c.Context.ResponseWriter, "ControllerName:"+c.Context.ControllerName)
	io.WriteString(c.Context.ResponseWriter, "ActionName:"+c.Context.ActionName)
	io.WriteString(c.Context.ResponseWriter, "name:"+c.Context.Request.Form.Get("name"))
}

func (c *UserController) UserInfo() {
	user := model.User{}
	rows, err := user.Select()
	if err != nil {
		c.Assign("name", err)
	} else {
		c.Assign("name", "唯一")
	}

	c.Assign("rows", rows)
	c.Display("")
}
