package cmd

import (
	"money-keeper/internal/database"
	"money-keeper/internal/server"

	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{
		Use: "money-keeper",
	}

	devCmd := &cobra.Command{
		Use: "dev",
		Run: func(cmd *cobra.Command, args []string) {
			database.Connect()
			server.RunServer()
		},
	}
	srvCmd := &cobra.Command{
		Use: "srv",
		Run: func(cmd *cobra.Command, args []string) {
			server.RunServer()
		},
	}
	dbCmd := &cobra.Command{
		Use: "db",
		Run: func(cmd *cobra.Command, args []string) {
			database.Connect()
			defer database.Close()
		},
	}
	migrateCmd := &cobra.Command{
		Use: "migrate",
	}
	migrateUpCmd := &cobra.Command{
		Use: "up",
		Run: func(cmd *cobra.Command, args []string) {
			MigrateUp()
		},
	}
	rootCmd.AddCommand(devCmd)
	rootCmd.AddCommand(srvCmd)
	rootCmd.AddCommand(dbCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.Execute()
}
