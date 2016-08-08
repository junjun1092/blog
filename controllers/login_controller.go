package main

import (
	"beego-blog-master/controllers"
	"beego-blog-master/g"
)

type LoginContorller struct {
	controllers.BaseController
}
func (this *LoginContorller) login(){
	this.TplNames = "login/login.html"
}
func (this *LoginContorller) DoLogin(){
	name := this.GetString("name")
	if(name == ""){
		this.Ctx.WriteString("姓名为空")
		return
	}
	password := this.GetString("password")
	if(password == ""){
		this.Ctx.WriteString("密码为空")
		return
	}
	this.Ctx.SetCookie("bb_name", name, 2592000, "/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "bb_password" + password + "; Max-Age=2592000; Path=/; httponly")
	this.Ctx.WriteString("")
}

func (this *LoginContorller) logout(){
	this.Ctx.WriteString("bb_name", g.RootName, 0, "/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "bb_password" + password + "; Max-Age=0; Path=/; httponly")
	this.Ctx.WriteString("logout")
}

