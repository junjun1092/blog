package global

import (
	"github.com/api.v7-develop/conf"
	"github.com/g"
)

var (
	RootEmail      string
	RootName       string
	RootPass       string
	RootPortrait   string
	BlogTitle      string
	BlogResume     string
	BlogLogo       string
	QiniuAccessKey string
	QiniuSecretKey string
	QiniuScope     string
	UseQiniu       bool
)

func initCfg() {
	RootName = g.Cfg.String("root_name")
	RootEmail = g.Cfg.String("root_email")
	RootPass = g.Cfg.String("root_pass")
	RootPortrait = g.Cfg.String("root_portrait")
	BlogTitle = g.Cfg.String("blog_title")
	BlogResume = g.Cfg.String("blog_resume")
	BlogLogo = g.Cfg.String("blog_logo")
	QiniuAccessKey = g.Cfg.String("qiniu_access_key")
	QiniuSecretKey = g.Cfg.String("qiniu_secret_key")
	QiniuScope = g.Cfg.String("qiniu_scope")
	UseQiniu = QiniuAccessKey != "" && QiniuSecretKey != "" && QiniuScope != ""
	conf.ACCESS_KEY = QiniuAccessKey
	conf.SECRET_KEY = QiniuSecretKey
}

func main() {
}
