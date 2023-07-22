package swagger

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func init() {
}

var Cmd = &cobra.Command{
	Use:   "swagger",
	Short: "生成接口 swagger 文件",
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

	},
}
