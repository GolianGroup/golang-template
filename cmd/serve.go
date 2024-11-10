package cmd

import (
	"context"
	"golang_template/app"
	"log"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve API",
	Run:   serve,
	Args:  cobra.MaximumNArgs(2),
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	log.Println("serve")
	application := app.NewApplication(context.TODO())
	application.Setup()
}
