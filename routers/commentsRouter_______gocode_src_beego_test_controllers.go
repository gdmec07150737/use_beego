package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego_test/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_test/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_test/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_test/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/user/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_test/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_test/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/user/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_test/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_test/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/user/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_test/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_test/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/users`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
