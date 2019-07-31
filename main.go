package main

import (
	_ "blog/models"
	_ "blog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Run()
}
