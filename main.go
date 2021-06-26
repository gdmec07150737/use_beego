package main

import (
	_ "beego_test/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run()
}

func init() {
	err := orm.RegisterDataBase("default", "mysql", "root:123456@/user_test?charset=utf8")
	if err != nil {
		panic("数据库链接失败:" + err.Error())
	}
}