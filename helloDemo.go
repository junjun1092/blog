package blog

import (
	"github.com/astaxie/beego"
)

type MainContorller struct {
	beego.Controller
}

func (this *MainContorller) Get() {
	this.Ctx.WriteString("hello world")
}
func main() {
	beego.Router("/", &MainContorller{})
	beego.Run()
}
