package controllers

import (
	"beego_test/models"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"*/

	var subject models.Subject
	err := func() error {
		id, err := c.GetInt64("id")
		beego.Info(id)
		if err != nil {
			id = 1
		}
		subject, err = models.GetSubjectById(id)
		if err != nil {
			return errors.New("subject not exist")
		}
		return nil
	}()
	if err != nil {
		c.Ctx.WriteString("数据不存在")
		return
	}
	var option map[string]string
	if err = json.Unmarshal([]byte(subject.Option), &option); err != nil {
		c.Ctx.WriteString("参数错误:不是json编码")
		return
	}
	c.Data["ID"] = subject.Id
	c.Data["Option"] = option
	c.Data["Img"] = "/static" + subject.Img
	httpPort := beego.BConfig.Listen.HTTPPort
	c.Data["HttpPort"] = httpPort
	c.TplName = "guess.html"
}

func (c *MainController) Post()  {
	var subject models.Subject
	err := func() error {
		id, err := c.GetInt64("id")
		beego.Info(id)
		if err != nil {
			id = 1
		}
		subject, err = models.GetSubjectById(id)
		if err != nil {
			return errors.New("subject not exist")
		}
		return nil
	}()
	if err != nil {
		c.Ctx.WriteString("wrong params")
	}
	answer := c.GetString("key")
	right := models.Answer(subject.Id, answer)
	c.Data["Right"] = right
	c.Data["Next"] = subject.Id + 1
	c.Data["ID"] = subject.Id
	httpPort := beego.BConfig.Listen.HTTPPort
	c.Data["HttpPort"] = httpPort
	c.TplName = "guess.html"
}