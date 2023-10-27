package main

import (
	"log"
	"odaLog/odaLog/database"
	"odaLog/odaLog/initializers"
	"odaLog/odaLog/models"
)

func init() {

	config, err := initializers.LoadConfig("..")
	if err != nil {
		log.Fatal("fatal error path from migrate", err)
	}
	database.ConnectDB(&config)
}

func main() {
	database.GlobalDB.AutoMigrate(models.Borcs{})
	database.GlobalDB.AutoMigrate(models.Users{})

}
