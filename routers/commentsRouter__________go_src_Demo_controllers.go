package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Demo/controllers:MainController"] = append(beego.GlobalControllerRouter["Demo/controllers:MainController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Demo/controllers:SqlController"] = append(beego.GlobalControllerRouter["Demo/controllers:SqlController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/sql`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
