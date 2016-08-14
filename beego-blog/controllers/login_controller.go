package controllers
import (
	"github.com/astaxie/beego"
	"fmt"
)
type LoginController struct{
	beego.Controller
}
func(this *LoginController) login(){
	this.TplName = "login/login.html"
}
func (this *LoginController) doLogin(){
	name := this.GetString("userName")
	fmt.Print("name:" + name)
	if name == "" {
		this.Ctx.WriteString("userName is empty")
		return
	}
	password:= this.GetString("password")
	if password == ""{
		this.Ctx.WriteString("password is empty")
		return
	}
	this.Ctx.SetCookie("bb_name",name,2592000,"/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie","bb_password=" + password + "; Max-Age=2592000;Path=/;httponly")
	this.Ctx.WriteString("")

}

