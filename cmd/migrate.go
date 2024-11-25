/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"golang_template/internal/config"
	"golang_template/internal/database/postgres"
	"log"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the migration files",
	Long: `Migrate the migration files. Example:
	migrate                              Migrate all the migration files
	migrate --dir ./database/migrations  Migrate all the migration files from sepecific directory
	migrate --version 1                  Migrate the migration file up to version 1
	migrate --fake true                  Fake apply all the migration files`,
	Run: func(cmd *cobra.Command, args []string) {
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

		fakeFlag, err := cmd.Flags().GetBool("fake")
		if err != nil {
			cmd.PrintErrf("Error while getting fake flag:\n\t %v", err)
			return
		}

		dbConfig, err := config.LoadConfig("config/config.yml")

		if err != nil {
			log.Fatalf("failed to setup viper: %s", err.Error())
			return
		}

		db, err := postgres.NewPostgresDatabase(cmd.Context(), &dbConfig.Postgres)
		if err != nil {
			cmd.PrintErrf("Error while initializing postgres database:\n\t %v", err)
			return
		}
		defer db.Close()

		migration := postgres.NewMigration(db.DB())
		err = migration.ApplyMigrations(dirFlag, versionFlag, fakeFlag)
		if err != nil {
			cmd.PrintErrf("Error while applying migrations:\n\t %v", err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().String("dir", "./internal/database/migrations", "Directory of the migrations")
	migrateCmd.Flags().Int64("version", 0, "Version of the migration that is going to be applied")
	migrateCmd.Flags().Bool("fake", false, "Fake apply migrations.")
}
