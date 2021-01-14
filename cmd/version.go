package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "unknown" //assigned at make
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version string of ddns-aliyun",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ddns-aliyun:" + version)
	},
}
