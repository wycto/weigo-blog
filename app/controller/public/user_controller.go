package public

import (
	"github.com/wycto/weigo"
	"github.com/wycto/weigo/datatype"
	"github.com/wycto/weigo/msg"
	"github.com/wycto/weigo/tools/datetime"
	"github.com/wycto/weigo/tools/encrypt"
	"weigo-blog/app/model"
)

type UserController struct {
	weigo.Controller
}

func (c *UserController) Login() {
	if c.Context.IsPost() {
		//参数校验
		if c.Context.NotHasOrEmpty("username", "password") {
			c.ResponseErrorMessage(msg.ParamsMissingORInvalid, nil)
			return
		}

		//存在校验
		row, err := model.UserModel().Where(&datatype.Row{"username": c.Context.Param("username")}).Find()
		if err != nil {
			c.ResponseErrorMessage(msg.DataAccountNotExists, nil)
			return
		}

		//密码校验
		if row.Get("password") != encrypt.MD5(c.Context.Param("password")) {
			c.ResponseErrorMessage(msg.DataPasswordError, nil)
			return
		}

		c.ResponseSuccess("登录成功", row)
	} else {
		c.MethodNotAllowed()
	}
	return
}

func (c *UserController) Register() {
	if c.Context.IsPost() {
		//参数校验
		if c.Context.NotHasOrEmpty("username", "password") {
			c.ResponseErrorMessage(msg.ParamsMissing, nil)
			return
		}

		//存在校验
		_, err := model.UserModel().Where(&datatype.Row{"username": c.Context.Param("username")}).Find()
		if err == nil {
			c.ResponseErrorMessage(msg.DataAccountExists, nil)
			return
		}

		//插入数据
		data := datatype.Row{"username": c.Context.Param("username")}
		data.Set("password", weigo.MD5(c.Context.Param("password")))
		data.Set("status", 1)
		data.Set("register_time", datetime.DateTime())
		data.Set("register_ip", weigo.GetIP())

		id, err := model.UserModel().Insert(&data)
		if err != nil {
			c.ResponseErrorMessage(msg.SysError, err.Error())
			return
		}

		c.ResponseSuccess("注册成功", id)
	} else {
		c.MethodNotAllowed()
	}
	return
}
