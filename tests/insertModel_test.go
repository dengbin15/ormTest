package test

import (
	"Demo/models"
	"github.com/astaxie/beego"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"testing"
)

func TestAddUser(t *testing.T) {
	/*user := UserInfo{Name:"dengbin",Age:12}
	id , err := AddUser(&user)
	fmt.Println(id,err)*/
	Context("testAddUser",func(){
		It("testAddUser",func(){
			user := models.UserInfo{Name:"dengbin",Age:12}
			err := models.AddUser(&user)
			beego.Info("Debug")
			g := gomega.NewGomegaWithT(t)
			g.Expect(err).To(gomega.Equal(nil))
			if err != nil {
				t.Fatal("fail!")
			}
		})
	})
}
