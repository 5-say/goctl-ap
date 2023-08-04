package tool

import (
	"embed"
	"fmt"
	"text/template"

	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func Generat(TPL embed.FS, data map[string]string, templatePath string, generatedFilePath string) {
	tpl, _ := template.ParseFS(TPL, templatePath)
	f, err := pathx.CreateIfNotExist(generatedFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	tpl.Execute(f, data)
}
