package service

import (
	"embed"

	"github.com/5-say/goctl-ap/tool"
)

//go:embed tpl/**/*
var TPL embed.FS

func generat(data map[string]string, templatePath string, generatedFilePath string) {
	tool.Generat(TPL, data, templatePath, generatedFilePath)
}
