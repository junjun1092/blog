package controllers

import (
	"github.com/astaxie/beego"
	"github.com/g"
)

type LoginContorller struct {
	beego.Controller
}


func (c *LoginContorller) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login/login.html"
}

func (this *LoginContorller) Login(){
	this.TplName = "login/login.html"
}

func (this *LoginContorller) DoLogin(){
	name := this.GetString("name")
	if name == "" {
		this.Ctx.WriteString("姓名为空")
		return 
	}
	password := this.GetString("password")
	if  password == ""{
		this.Ctx.WriteString("密码为空")
		return 
	}
	if g.RootName != name {
		this.Ctx.WriteString("用户名不正确")
		return
	}
	if g.RootPass != password {
		this.Ctx.WriteString("密码不正确")
		return
	}
	this.Ctx.SetCookie("bb_name", name ,2592000, "/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie","bb_password" + password + ";Max-Age=259200;Path=/;httponly" )
	this.Ctx.WriteString("")
}
func (this *LoginContorller) logout() {
	this.Ctx.SetCookie("bb_name",g.RootName,0,"/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie","bb_password=" + g.RootPass + ";Max-Age=0;Path=/;httponly")
	this.Ctx.WriteString("logout")
}