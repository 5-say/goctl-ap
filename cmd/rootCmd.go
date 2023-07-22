package cmd

import (
	"fmt"
	"os"

	"github.com/5-say/goctl-ap/cmd/swagger"
	"github.com/5-say/goctl-ap/cmd/ts"

	"github.com/spf13/cobra"
)

// Execute ..
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "ap",
	Short: "goctl api plugin",
	// Run:   rootRun,
}

func init() {
	// 参数
	// rootCmd.Flags().String("api", "", "The api file")
	// rootCmd.Flags().String("dir", "", "The target dir")

	rootCmd.AddCommand(ts.Cmd)
	rootCmd.AddCommand(swagger.Cmd)
}

func rootRun(cmd *cobra.Command, args []string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("Error: ", err)
		}
	}()

}
