package models

import (
	"Demo/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

type UserInfo struct{
	ClusterId int32   				`json:"id" orm:"pk;size(36);column(my_id)"`
	Name string				`json:"name"`
	Age int16				`orm:"column(my_age)"`
}

type SqlController struct {
	beego.Controller
}

// init()前面没有(c *SqlController)
// init()函数没有接受参数和返回值
/*func init(){
	orm.Debug = true
	/*
	 * 出错点：
	 * 	1.charset=utf8而不是chacter=utf8
	 *  2.test后面是?,不是@
	 *  3.要先注册UserInfo才能newOrm()
	 */
/*
	err := orm.RegisterDataBase("default","mysql","root:root@tcp(127.0.0.1:3306)/test?charset=utf8",30)
	if err != nil {
		fmt.Printf("error register database. %s",err.Error())
	}
	//orm.RegisterModel(new(UserInfo))
	orm.RegisterModel(new(UserInfo))
	orm.RunSyncdb("default",false,true)
	logs.Trace("Init done!")
}*/

func AddUser(user *UserInfo) error{
	err := orm.RegisterDataBase("default","mysql","root:root@tcp(127.0.0.1:3306)/test?charset=utf8",30)
	if err != nil {
		fmt.Printf("error register database. %s",err.Error())
	}
	//orm.RegisterModel(new(UserInfo))
	orm.RegisterModel(new(UserInfo))
	orm.RunSyncdb("default",false,true)
	logs.Trace("Init done!")
	log := logs.NewLogger(100)
	/*
	 * filename:logtest 保存的文件名
	 * maxlines 每个文件保存的最大行数
	 * maxsize 每个文件保存的最大尺寸
	 * daily 是否按照每天logrotate，默认是true
	 */
	// "filename":"D:\test.log\"
	//log.SetLogger("file",map)
	config := make(map[string]interface{})
	// config["filename"] = "e:/golang/go_pro/logs/logcollect.log"
	// config["filename"] = "e://golang//go_pro//logs//logcollect.log"
	config["filename"] = "d://test.log"
	config["level"] = logs.LevelDebug

	configStr,_ := json.Marshal(config)

	logs.SetLogger(logs.AdapterFile, string(configStr))

	log.Trace("this is log.Trace")

	//log.Trace("this is log.Trace")
	//beego.SetLogger("file",'{"filename":"logs/test.log"}')

	//beego.SetLevel(beego.LevelInfo)

	//beego.SetLevel(beego.LevelDebug)

	o = orm.NewOrm()
	o.Using("default")
	//fmt.Println("enter AddUser,going to enter Insert()")
	// 易错点：这里的user前面不需要加"&"
	//_,error := o.Insert(user)
	//fmt.Println("Insert() done!")

	var userInfo models.UserInfo
	qs :=o.QueryTable("user_info")
	err =qs.Filter("id",1).Filter("name","dengbin").One(&userInfo)
	if err != nil{
		beego.Error("query user_info failed")
	}
	return err
}

func (u *UserInfo) TableName() string{
	return "my_info"
}