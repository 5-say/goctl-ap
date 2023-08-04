package service

import (
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func genAPICode(c Config) {

	var (
		appService = c.appDir + "/" + c.serviceName
		data       = map[string]string{
			"serviceName": c.serviceName,
			"apiPort":     c.apiPort,
			"docPort":     c.docPort,
		}
	)

	// 创建目录
	pathx.MkdirIfNotExist(appService + "/api/doc")
	pathx.MkdirIfNotExist(appService + "/api/etc")
	pathx.MkdirIfNotExist(appService + "/api/internal/config")
	pathx.MkdirIfNotExist(appService + "/api/internal/logic/api")
	pathx.MkdirIfNotExist(appService + "/api/internal/middleware")
	pathx.MkdirIfNotExist(appService + "/api/internal/svc")
	pathx.MkdirIfNotExist(appService + "/db/dao")

	// 生成文件
	generat(data, "tpl/app/api/doc/swagger.go.tpl", appService+"/api/doc/swagger.go")
	generat(data, "tpl/app/api/doc/swagger.json.tpl", appService+"/api/doc/swagger.json")
	generat(data, "tpl/app/api/etc/example.yaml.tpl", appService+"/api/etc/example.yaml")
	generat(data, "tpl/app/api/etc/service.yaml.tpl", appService+"/api/etc/"+c.serviceName+".yaml")
	generat(data, "tpl/app/api/internal/config/config.go.tpl", appService+"/api/internal/config/config.go")
	generat(data, "tpl/app/api/internal/middleware/authadminMiddleware.go.tpl", appService+"/api/internal/middleware/authadminMiddleware.go")
	generat(data, "tpl/app/api/internal/svc/serviceContext.go.tpl", appService+"/api/internal/svc/serviceContext.go")
	generat(data, "tpl/app/api/service.go.tpl", appService+"/api/"+c.serviceName+".go")
	generat(data, "tpl/app/db/db.go.tpl", appService+"/db/db.go")
	generat(data, "tpl/app/db/dao/dao.go.tpl", appService+"/db/dao/dao.go")
	generat(data, "tpl/app/.gitignore.tpl", appService+"/.gitignore")

}
