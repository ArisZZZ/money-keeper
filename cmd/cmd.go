package cmd

import "money-keeper/internal/database"

func RunServer() {
	database.Connect()
}
