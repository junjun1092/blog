package controllers

import (
	"github.com/astaxie/beego"
)

type LoginContorller struct {
	beego.Controller
}

func (c *LoginContorller) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login.html"
}
