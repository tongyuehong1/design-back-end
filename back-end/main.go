package main

import (
	_ "github.com/tongyuehong1/design-back-end/back-end/routers"
	"github.com/astaxie/beego"
	mysql "github.com/tongyuehong1/design-back-end/back-end/init"
)

func main() {
	mysql.InitSql()
	beego.Run()
}

