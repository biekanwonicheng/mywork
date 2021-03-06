package main

import (
	_ "mywork/routers"
	"github.com/astaxie/beego"
	"strings"
	_"mywork/models"

	"encoding/gob"
	"mywork/models"
	"strconv"
)

func main() {

	initSetSession()

	setTemplate()

	gob.Register(models.LiteOauthUser{})

	beego.Run()
}


func setTemplate()  {

	//创建新的模板函数
	beego.AddFuncMap("equrl", func(x,y string)bool {

		//去除获取url中的 /
		x1 := strings.Trim(x,"/")

		y1 := strings.Trim(y,"/")

		if strings.Compare(x1,y1) == 0 {

			return true
		}else{

			return false
		}

	})


	//拼接参数
	beego.AddFuncMap("urljoin", func(x string,y int) string {

		str := strconv.Itoa(y)
		return x+str
	})

	/**
		判读一个参数是否存在另一个集合里面
	 */
	 beego.AddFuncMap("in_array", func(x uint,sli []string) bool {

	 	str := strconv.Itoa(int(x))

	 	for _,val := range sli{

	 		if val == str{

	 			return true
	 			break
			}
		}

		return false
	 })

}


/**
	初始化session
 */
func initSetSession()  {

	beego.BConfig.WebConfig.Session.SessionOn = true

	//设置 cookies 的名字，Session 默认是保存在用户的浏览器 cookies 里面的，默认名是 beegosessionID，配置文件对应的参数名是
	beego.BConfig.WebConfig.Session.SessionName = "liteblog"

	beego.BConfig.WebConfig.Session.SessionProvider="file"

	beego.BConfig.WebConfig.Session.SessionProviderConfig = "docs/session"
}
