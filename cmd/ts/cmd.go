package ts

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func init() {
}

var Cmd = &cobra.Command{
	Use:   "ts",
	Short: "生成前端 TypeScript 接口文件",
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

		// 构造请求包
		if err := genComponents(p.Dir, p.Api); err != nil {
			panic(err)
		}
		if err := genAPI(p.Dir, p.Api); err != nil {
			panic(err)
		}
	},
}
