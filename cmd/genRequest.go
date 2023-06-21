package cmd

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

//go:embed tpl/apiRequest.ts
var requestTemplate string

func genRequest(dir string) error {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	filename := filepath.Join(abs, "apiRequest.ts")
	if pathx.FileExists(filename) {
		return nil
	}

	return os.WriteFile(filename, []byte(requestTemplate), 0644)
}
