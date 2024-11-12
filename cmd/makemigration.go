/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"golang_template/internal/database"
	"golang_template/internal/utils"
	"log"

	"github.com/spf13/cobra"
)

// makemigrationCmd represents the makemigration command
var makemigrationCmd = &cobra.Command{
	Use:   "makemigration [name]",
	Short: "Create a new migration file",
	Long: `Create a new migration file. Example:
		makemigration add_users_table // Create a new migration file with sql type
		makemigration add_users_table --type go // Create a new migration file with go type
		makemigration add_users_table --dir ./database/migrations // Create a new migration file in specific directory`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("makemigration")
		name := args[0]
		if name == "" {
			cmd.PrintErrf("Migration name is required")
			return
		}

		typeFlag, err := cmd.Flags().GetString("type")
		if err != nil {
			cmd.PrintErrf("Error while getting type flag:\n\t %v", err)
			return
		}
		if typeFlag != "sql" && typeFlag != "go" {
			cmd.PrintErrf("Invalid migration type: %s. Must be 'sql' or 'go'", typeFlag)
			return
		}

		dirFlag, err := cmd.Flags().GetString("dir")
		if err != nil {
			cmd.PrintErrf("Error while getting dir flag:\n\t %v", err)
			return
		}

		config, err := utils.SetupViper("config/config.yml")

		if err != nil {
			log.Fatalf("failed to setup viper: %s", err.Error())
			return
		}

		db, err := database.NewDatabase(cmd.Context(), &config.DB)
		if err != nil {
			cmd.PrintErrf("Error while initializing database:\n\t %v", err)
			return
		}
		defer db.Close()

		migration := database.NewMigration(db.DB())
		err = migration.CreateMigrationFile(dirFlag, name, typeFlag)
		if err != nil {
			cmd.PrintErrf("Error while creating migration file:\n\t %v", err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(makemigrationCmd)

	makemigrationCmd.Flags().String("dir", "./internal/database/migrations", "Directory of the migrations")
	makemigrationCmd.Flags().String("type", "sql", "Type of the migration file which is sql or go")
}
