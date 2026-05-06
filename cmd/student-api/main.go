package main

import (
	"example/rest-api/config"
	"example/rest-api/database"
	"fmt"
)

func main() {
	fmt.Println("welcome")
	config.Load()

	dbUser := config.AppConfig.Database.Username
	database.Connect()
	fmt.Println("Database User:", dbUser)

}
