package models

import (
	"errors"
	"github.com/agiledragon/gomonkey"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"testing"
)

func TestAddUser2(t *testing.T) {
	/*user := UserInfo{Name:"dengbin",Age:12}
	id , err := AddUser(&user)
	fmt.Println(id,err)*/
	Context("testAddUser",func(){
		It("testAddUser",func(){
			user := UserInfo{Name:"dengbin",Age:12}
			err := AddUser(&user)
			beego.Info("Debug")
			g := gomega.NewGomegaWithT(t)
			g.Expect(err).To(gomega.Equal(nil))
			if err != nil {
				t.Fatal("fail!")
			}
		})
	})
}

func TestAddUser(t *testing.T){
	patch := gomonkey.ApplyFunc(orm.RegisterDataBase,func() error{
		err := errors.New("error!")
		return err
	})
	defer patch.Reset()
	user := UserInfo{Name:"dengbin",Age:12}
	err := AddUser(&user)
	if err != nil {
		t.Fatal("fail!")
	}
}


