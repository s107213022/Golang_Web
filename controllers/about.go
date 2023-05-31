package controllers

import (
	"github.com/astaxie/beego"
)

// AboutController operations for About
type AboutController struct {
	beego.Controller
}

// URLMapping ...
func (c *AboutController) URLMapping() {
	c.Mapping("Get", c.Get)
}


// Get
// @Title GetOne
// @Description get About by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.About
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AboutController) Get() {
	// View (Template) 裡要用到的變數
    c.Data["Name"] = "John Doe"
    // 指定 View (Template) 檔名
    c.TplName = "about.tpl"
}


