package index

import (
	"fmt"
	"github.com/wycto/weigo"
	"weigo-blog/app/model"
)

type ModelController struct {
	weigo.Controller
}

func (receiver *ModelController) GetOne() {
	row, _ := model.UserModel().Find()
	fmt.Println(row)

	fmt.Println(model.UserModel().Test())
}
