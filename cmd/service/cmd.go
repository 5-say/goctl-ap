package service

import (
	"fmt"
	"runtime"

	"github.com/5-say/goctl-ap/tool"
	"github.com/spf13/cobra"
)

func init() {
	// 参数
	Cmd.Flags().StringP("app", "a", "", "The app directory")
	Cmd.Flags().StringP("define", "d", "", "The define directory")
}

// Config ..
type Config struct {
	appDir      string // 代码生成目录
	defineDir   string // 设计定义目录
	serviceName string // 服务名
	apiPort     string // 启动 api 服务的端口号
	docPort     string // 启动 文档 服务的端口号
	// rpcPort     string // 启动 rpc 服务的端口号
}

var Cmd = &cobra.Command{
	Use:   "service",
	Short: "生成服务初始化文件",
	Long:  "独立交互功能，不依赖 goctl 启动",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Error: ", err)
				for i := 0; i < 10; i++ {
					_, file, line, _ := runtime.Caller(i)
					fmt.Printf("%s:%d\n", file, line)
				}
			}
		}()

		c := Config{}
		c.appDir = tool.GetFlagString(cmd, "app")
		c.defineDir = tool.GetFlagString(cmd, "define")

		for c.serviceName == "" {
			fmt.Print("Please enter service name:")
			fmt.Scanln(&c.serviceName)
		}

		for c.apiPort == "" {
			fmt.Print("Please enter api port:")
			fmt.Scanln(&c.apiPort)
		}

		for c.docPort == "" {
			fmt.Print("Please enter doc port:")
			fmt.Scanln(&c.docPort)
		}

		// fmt.Print("Please enter rpc port:")
		// fmt.Scanln(&c.rpcPort)

		// 构造
		genDefineFile(c)
		genAPICode(c)
	},
}
