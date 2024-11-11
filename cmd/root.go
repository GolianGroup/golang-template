package cmd

import (
	"github.com/spf13/cobra"
	"master/internal/di"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "template",
	Short: "All of template's commands",
}

func Execute() {
	di.Start()
	//err := RootCmd.Execute()
	//if err != nil {
	//	os.Exit(1)
	//}
}
