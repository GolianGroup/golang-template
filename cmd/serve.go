package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"master/internal/di"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve API",
	Run:   Serve,
	Args:  cobra.MaximumNArgs(2),
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func Serve(cmd *cobra.Command, args []string) {
	log.Println("serve")
	di.Start()
}
