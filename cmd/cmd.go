package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

// Execute ..
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "ts",
	Short: "Generate ts files for provided api in api file",
	Run:   rootRun,
}

func init() {
	// 参数
	// rootCmd.Flags().String("api", "", "The api file")
	// rootCmd.Flags().String("dir", "", "The target dir")
}

func rootRun(cmd *cobra.Command, args []string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("Error: ", err)
		}
	}()

	// 获取 goctl api 上下文
	p, err := plugin.NewPlugin()
	if err != nil {
		panic(err)
	}

	// 构造请求包
	if err := genRequester(p.Dir); err != nil {
		panic(err)
	}
	if err := genComponents(p.Dir, p.Api); err != nil {
		panic(err)
	}
	if err := genAPI(p.Dir, p.Api); err != nil {
		panic(err)
	}
}
