package main

import (
	"log"
	"odaLog/odaLog/database"
	"odaLog/odaLog/initializers"
	"odaLog/odaLog/router"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("configiration is not defined")
	}

	database.ConnectDB(&config)

}

func main() {
	r := router.RouterConfig()
	r.Run()
}
