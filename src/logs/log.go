package logs

import (
	io "io/ioutil"
	"github.com/astaxie/beego"
	"os"
)

func init() {
	filename := beego.AppConfig.String("logfile")
	pwd, _ := os.Getwd()
	data, err := io.ReadFile(pwd + "/" + filename)
	if err != nil {
		beego.Warn("No Log Config File!")
		return
	}

	beego.SetLogger("file", string(data))
	//beego.BeeLogger.DelLogger("console")
	beego.SetLogFuncCall(true)
}
