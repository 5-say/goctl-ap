package tool

import "github.com/spf13/cobra"

func GetFlagString(cmd *cobra.Command, name string) string {
	flag := cmd.Flag(name)
	if flag != nil {
		return flag.Value.String()
	}
	return ""
}
