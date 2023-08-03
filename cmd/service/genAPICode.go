package service

import (
	"fmt"
	"text/template"

	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func genAPICode(c Config) {

	var (
		tplPath     = "cmd/service/tpl/app"
		servicePath = c.appDir + "/" + c.serviceName
	)

	// 创建目录
	pathx.MkdirIfNotExist(servicePath + "/api/doc")
	pathx.MkdirIfNotExist(servicePath + "/api/etc")
	pathx.MkdirIfNotExist(servicePath + "/api/internal/config")
	pathx.MkdirIfNotExist(servicePath + "/api/internal/logic/api")
	pathx.MkdirIfNotExist(servicePath + "/api/internal/middleware")
	pathx.MkdirIfNotExist(servicePath + "/api/internal/svc")
	pathx.MkdirIfNotExist(servicePath + "/db/dao")

	// 创建 .gitignore
	{
		tpl, _ := template.ParseFiles(tplPath + "/.gitignore.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/.gitignore")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, nil)
	}

	// 创建 api/doc/swagger.go
	{
		tpl, _ := template.ParseFiles(tplPath + "/api/doc/swagger.go.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/api/doc/swagger.go")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, nil)
	}

	// 创建 api/doc/swagger.json
	{
		tpl, _ := template.ParseFiles(tplPath + "/api/doc/swagger.json.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/api/doc/swagger.json")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, nil)
	}

	// 创建 api/etc/example.yaml
	{
		tpl, _ := template.ParseFiles(tplPath + "/api/etc/example.yaml.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/api/etc/example.yaml")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, map[string]string{
			"serviceName": c.serviceName,
			"apiPort":     c.apiPort,
			"docPort":     c.docPort,
		})
	}

	// 创建 api/etc/service.yaml
	{
		tpl, _ := template.ParseFiles(tplPath + "/api/etc/service.yaml.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/api/etc/" + c.serviceName + ".yaml")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, map[string]string{
			"serviceName": c.serviceName,
			"apiPort":     c.apiPort,
			"docPort":     c.docPort,
		})
	}

	// 创建 api/internal/config/config.go
	{
		tpl, _ := template.ParseFiles(tplPath + "/api/internal/config/config.go.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/api/internal/config/config.go")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, nil)
	}

	// 创建 api/internal/middleware/authadminMiddleware.go
	{
		tpl, _ := template.ParseFiles(tplPath + "/api/internal/middleware/authadminMiddleware.go.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/api/internal/middleware/authadminMiddleware.go")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, nil)
	}

	// 创建 api/internal/svc/serviceContext.go
	{
		tpl, _ := template.ParseFiles(tplPath + "/api/internal/svc/serviceContext.go.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/api/internal/svc/serviceContext.go")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, map[string]string{
			"serviceName": c.serviceName,
		})
	}

	// 创建 api/service.go
	{
		tpl, _ := template.ParseFiles(tplPath + "/api/service.go.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/api/" + c.serviceName + ".go")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, map[string]string{
			"serviceName": c.serviceName,
		})
	}

	// 创建 db/db.go
	{
		tpl, _ := template.ParseFiles(tplPath + "/db/db.go.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/db/db.go")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, map[string]string{
			"serviceName": c.serviceName,
		})
	}

	// 创建 db/dao/dao.go
	{
		tpl, _ := template.ParseFiles(tplPath + "/db/dao/dao.go.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/db/dao/dao.go")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, map[string]string{
			"serviceName": c.serviceName,
		})
	}

}
