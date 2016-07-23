package blog

import (
	"fmt"
	"github.com/astaxie/beego"
)

func (this *MainController) blog() {
	name, _ := this.getString(":name")
	password := this.getString("password")
	beego.Debug(fmt.Sprintf("Catelog name: %d password:%v", name, password))
}

func main() {
	beego.Router("/blog/:name/:password", &controller{}, "get:blog")
}
