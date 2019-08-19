package models

import (
	_ "gitserver/kubernetes/dockerstack-cluster/routers"
	"path/filepath"
	"runtime"

	"github.com/astaxie/beego"
)

func init() {

	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.BConfig.CopyRequestBody = true
	beego.TestBeegoInit(apppath)

}
