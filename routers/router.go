package routers

import (
  "Demo/controllers"
  "github.com/astaxie/beego"
)

func init() {
  beego.Include(&controllers.MainController{})
  //beego.Router("/sql",&controllers.SqlController{},"get:Get")
  beego.Include(&controllers.SqlController{})

/*  ns :=
    beego.NewNamespace("/v1",
      beego.NSCond(func(ctx *context.Context) bool {
        if ctx.Input.Domain() == "api.beego.me"{
          return true
        }
        return false
      })；
    )*/
}
