package cmd

import (
	"money-keeper/internal/database"
	"money-keeper/internal/server"

	"github.com/spf13/cobra"
)

var (
	step int
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
	MigrateCreateCmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			CreateMigrate(args[0])
		},
	}
	migrateUpCmd := &cobra.Command{
		Use: "up",
		Run: func(cmd *cobra.Command, args []string) {
			MigrateUp()
		},
	}
	migrateStepsCmd := &cobra.Command{
		Use: "steps",
		Run: func(cmd *cobra.Command, args []string) {
			MigrateSteps(step)
		},
	}
	migrateStepsCmd.Flags().IntVar(&step, "int", -1, "Set step")
	migrateCmd.AddCommand(migrateUpCmd, MigrateCreateCmd, migrateStepsCmd)
	rootCmd.AddCommand(devCmd, srvCmd, dbCmd, migrateCmd)
	rootCmd.Execute()
}
