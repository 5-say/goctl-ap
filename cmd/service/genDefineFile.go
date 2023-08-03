package service

import (
	"fmt"
	"text/template"

	"github.com/jinzhu/inflection"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func genDefineFile(c Config) {

	var (
		tplPath     = "cmd/service/tpl/define"
		servicePath = c.defineDir + "/" + c.serviceName
	)

	// 创建目录
	pathx.MkdirIfNotExist(servicePath + "/api/app")
	pathx.MkdirIfNotExist(servicePath + "/api/platform")
	pathx.MkdirIfNotExist(servicePath + "/db/mysql")

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

	// 创建 Makefile
	{
		tpl, _ := template.ParseFiles(tplPath + "/Makefile.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/Makefile")
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

	// 创建 db/mysql/demos.sql
	{
		tableName := inflection.Plural(c.serviceName)
		tpl, _ := template.ParseFiles(tplPath + "/db/mysql/demo.sql.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/db/mysql/" + tableName + ".sql")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, map[string]string{
			"tableName": tableName,
		})
	}

	// 创建 api/manager_admin.api.tpl
	{
		tpl, _ := template.ParseFiles(tplPath + "/api/manager_admin.api.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/api/manager_admin.api")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, map[string]string{
			"serviceName": c.serviceName,
		})
	}

	// 创建 api/app/app.api.tpl
	{
		tpl, _ := template.ParseFiles(tplPath + "/api/app/app.api.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/api/app/app.api")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, map[string]string{
			"serviceName": c.serviceName,
		})
	}

	// 创建 api/platform/manager.api
	{
		tpl, _ := template.ParseFiles(tplPath + "/api/platform/manager.api.tpl")

		f, err := pathx.CreateIfNotExist(servicePath + "/api/platform/manager.api")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		tpl.Execute(f, nil)
	}
}
