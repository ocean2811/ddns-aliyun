package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var gDebug bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&gDebug, "debug", "", false, "show debug log")
	// rootCmd.Flags().BoolVarP(&gDebug, "debug", "", false, "show debug log")
}

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(">>dns-aliyun<<")
		fmt.Println("Use \"dns-aliyun --help\" for more information about the application.")
	},
}

// Execute cmd parse
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if gDebug {
			fmt.Printf("%+v\n", err)
		}
	}
}
