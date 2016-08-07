package main

import "beego-blog-master/controllers"

type LoginContorller struct {
	controllers.BaseController
}
func (this *LoginContorller) login()
{
	this.TplNames = "login/login.html"
}
func (this *LoginContorller) DoLogin()
{
	name := this.GetString("name")
	if(name == "")
	{
		this.Ctx.WriteString("姓名为空")
		return
	}
	password := this.GetString("password")
	if(password == "")
	{
		this.Ctx.WriteString("密码为空")
		return
	}
}
func main() {
}
