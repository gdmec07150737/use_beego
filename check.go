package main

import (
	"database/sql"
	"errors"
	"github.com/astaxie/beego/toolbox"
)

type DatabaseCheck struct {
}

func (* DatabaseCheck) Check() error {
	Db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/user_test?charset=utf8")
	defer Db.Close()
	if err != nil {
		return errors.New("数据源的名称不合法")
	}
	err = Db.Ping()
	if err != nil {
		return errors.New("数据库链接失败:" + err.Error())
	}
	return nil
}

func init()  {
	toolbox.AddHealthCheck("database", &DatabaseCheck{})
}