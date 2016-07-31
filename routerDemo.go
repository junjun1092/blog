package blog

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainContorller struct {
	beego.Controller
}

func (this *MainContorller) Blog() {
	catId, _ := this.GetInt(":catId")
	catName := this.GetString("catName")
	catPublish, _ := this.GetBool("catPublish")
	catPrice, _ := this.GetFloat(":catPrice")
	beego.Debug(fmt.Sprintf("Category Id:%d Name:%s Publish:%v Price:%v Price:%f", catId, catName, catPublish, catPrice))
}

func main() {
	beego.Router("blog/:catId/:catName:/:catPublish:/:catPrice", &MainContorller{}, "get:Blog")
}
