package cmd

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed tpl/apiRequester.ts
var requesterTemplate string

func genRequester(dir string) error {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	filename := filepath.Join(abs, "apiRequester.ts")
	// if pathx.FileExists(filename) {
	// 	return nil
	// }

	return os.WriteFile(filename, []byte(requesterTemplate), 0644)
}
