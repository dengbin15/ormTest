package controllers

import (
	"Demo/models"
	"fmt"
	"github.com/astaxie/beego"
)

type SqlController struct{
	beego.Controller
}

// @router /sql [get]
func (c *SqlController) Get(){
	// 易错点：Name,Age后面的值不能有空格
	user := models.UserInfo{Name:"dengbin",Age:12}
	//fmt.Println("going to enter AddUser")
	err := models.AddUser(&user)
	c.Ctx.WriteString(fmt.Sprintf("error : %v",err))
}
