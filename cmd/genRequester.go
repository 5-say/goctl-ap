package cmd

import (
	_ "embed"
	"os"
	"path/filepath"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

//go:embed tpl/apiRequester.ts
var requesterTemplate string

func genRequester(dir string, api *spec.ApiSpec) error {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	filename := strings.Replace(api.Service.Name, "-api", "", 1) + "Requester.ts"
	filename = filepath.Join(abs, filename)
	if pathx.FileExists(filename) {
		return nil
	}

	return os.WriteFile(filename, []byte(requesterTemplate), 0644)
}
