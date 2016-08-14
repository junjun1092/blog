package controllers

import (
	"github.com/astaxie/beego"
)

type Checker interface {
	checkLogin()
}
type BaseController struct {
	beego.Controller
	IsAdmin bool
}

func (this *BaseController) Prepare() {

	this.AssigonIsAdmin()

}

func (this *BaseController) AssigonIsAdmin() {
	bb_name := this.Ctx.GetCookie("bb_name")
	bb_password := this.Ctx.GetCookie("bb_password")
	if bb_name == "" || bb_password == "" {
		this.IsAdmin = false
		return
	}
	//if bb_name != g.RootName || bb_password != g.RootPass {
	//	this.IsAdmin = false
	//}
	this.IsAdmin = true
	this.Data["isAdmin"] = this.IsAdmin
}

