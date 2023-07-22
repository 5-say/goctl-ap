package swagger

import (
	"fmt"

	"github.com/5-say/goctl-ap/tool"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func init() {
	// 参数
	Cmd.Flags().StringP("filename", "f", "swagger.json", "The swagger file")
	Cmd.Flags().String("host", "", "The api host")
	Cmd.Flags().String("base", "", "The api base path")
}

var Cmd = &cobra.Command{
	Use:   "swagger",
	Short: "生成 swagger json 文件",
	Run: func(cmd *cobra.Command, args []string) {
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

		// 自动创建输出目录
		pathx.MkdirIfNotExist(p.Dir)

		var (
			filename = tool.GetFlagString(cmd, "filename")
			host     = tool.GetFlagString(cmd, "host")
			basePath = tool.GetFlagString(cmd, "basePath")
		)
		if err := Generate(filename, host, basePath, p); err != nil {
			panic(err)
		}
	},
}
