package service

import (
	_ "embed"

	"github.com/jinzhu/inflection"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func genDefineFile(c Config) {

	var (
		defineService = c.defineDir + "/" + c.serviceName
		data          = map[string]string{
			"serviceName": c.serviceName,
			"apiPort":     c.apiPort,
			"docPort":     c.docPort,
			"tableName":   inflection.Plural(c.serviceName),
		}
	)

	// 创建目录
	pathx.MkdirIfNotExist(defineService + "/api/app")
	pathx.MkdirIfNotExist(defineService + "/api/platform")
	pathx.MkdirIfNotExist(defineService + "/db/mysql")

	// 生成文件
	generat(data, "tpl/define/api/app/app.api.tpl", defineService+"/api/app/app.api")
	generat(data, "tpl/define/api/platform/manager.api.tpl", defineService+"/api/platform/manager.api")
	generat(data, "tpl/define/api/manager_admin.api.tpl", defineService+"/api/manager_admin.api")
	generat(data, "tpl/define/db/mysql/demo.sql.tpl", defineService+"/db/mysql/"+data["tableName"]+".sql")
	generat(data, "tpl/define/.gitignore.tpl", defineService+"/.gitignore")
	generat(data, "tpl/define/Makefile.tpl", defineService+"/Makefile")

}
