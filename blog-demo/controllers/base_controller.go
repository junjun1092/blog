package controllers

import (
	"github.com/astaxie/beego"
	"github.com/g"
)

type Checker interface {
	checkLogin()
}
type BaseController struct {
	beego.Controller
	IsAdmin bool
}

func (this *BaseController) Prepare(){
	this.Data["BlogLogo"] = g.BlogLogo
	this.Data["BlogTitle"] = g.BlogTitle
	this.Data["BlogResume"] = g.BlogResume
	this.Data["RootName"] = g.RootName
	this.Data["RootEmail"] = g.RootEmail
	this.Data["RootPortrait"] = g.RootPortrait
	this.AssigonIsAdmin()

}

func (this *BaseController) AssigonIsAdmin(){
	bb_name := this.Ctx.GetCookie("bb_name")
	bb_password := this.Ctx.GetCookie("bb_password")
	if bb_name == "" || bb_password == "" {
		this.IsAdmin = false
		return
	}
	if bb_name != g.RootName  || bb_password != g.RootPass{
		this.IsAdmin = false
	}
	this.IsAdmin = true
	this.Data["isAdmin"] = this.IsAdmin
}
func main() {
}
