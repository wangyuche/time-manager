package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(gettimeCmd)
	RootCmd.AddCommand(settimeCmd)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "time-manager",
	Short: "time-manager",
	Long:  `time-manager controls the time server manager.`,
}

var gettimeCmd = &cobra.Command{
	Use:   "get-time",
	Short: "get time server's time",
	Long:  `get time server's time`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get-time")
	},
}

var settimeCmd = &cobra.Command{
	Use:   "set-time",
	Short: "set time server's time",
	Long:  `set time server's time`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set-time")
	},
}
