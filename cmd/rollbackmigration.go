/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"golang_template/internal/config"
	"golang_template/internal/database"
	"golang_template/internal/database/postgres"
	"log"

	"github.com/spf13/cobra"
)

// rollbackmigrationCmd represents the rollbackmigration command
var rollbackmigrationCmd = &cobra.Command{
	Use:   "rollbackmigration",
	Short: "Rollback migration/migrations using this command",
	Long: `Rollback migration file/files using goose. The path should be path to migration files.
		You can write the migration hash to rollback to a specific version.
		For example:
		rollbackmigration
		rollbackmigration --dir ./database/migrations --version 1`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("rollbackmigration")

		versionFlag, err := cmd.Flags().GetInt64("version")
		if err != nil {
			cmd.PrintErrf("Error while getting version flag:\n\t %v", err)
			return
		}

		dirFlag, err := cmd.Flags().GetString("dir")
		if err != nil {
			cmd.PrintErrf("Error while getting dir flag:\n\t %v", err)
			return
		}

		dbConfig, err := config.LoadConfig("config/config.yml")

		if err != nil {
			log.Fatalf("failed to setup viper: %s", err.Error())
			return
		}

		db, err := postgres.NewDatabase(cmd.Context(), &dbConfig.DB)
		if err != nil {
			cmd.PrintErrf("Error while initializing database:\n\t %v", err)
			return
		}
		defer db.Close()

		migration := database.NewMigration(db.DB())
		err = migration.RollbackMigrations(dirFlag, versionFlag)
		if err != nil {
			cmd.PrintErrf("Error while rolling back migrations:\n\t %v", err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(rollbackmigrationCmd)

	rollbackmigrationCmd.Flags().String("dir", "./internal/database/migrations", "Directory of the migrations")
	rollbackmigrationCmd.Flags().Int64("version", 0, "Version of the migration that migrations will be rolled back to")
}
