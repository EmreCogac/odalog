package controllers

import (
	"log"
	"odaLog/odaLog/database"
	"odaLog/odaLog/initializers"
	"odaLog/odaLog/models"

	"github.com/gin-gonic/gin"
)

func init() {

	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("controllers is not working  load config")
	}

	database.ConnectDB(&config)

}

func GetUsers(c *gin.Context) {
	db := database.GlobalDB
	Users := []models.Users{}

	err := db.Find(&Users)

	if err != nil {
		c.JSON(500, gin.H{
			"users not found": err,
		})
		c.Abort()
	}
	c.JSON(200, gin.H{
		"Users": Users,
	})
}

func GetUserBorcs(c *gin.Context) {
	db := database.GlobalDB
	Borc := []models.Borcs{}

	err := db.Preload("Users").Find(&Borc)
	if err != nil {
		c.JSON(500, gin.H{
			"have some issues": err,
		})
	}

	c.JSON(200, gin.H{
		"user preload": Borc,
	})
}

func GetUsersFromName(c *gin.Context) {
	db := database.GlobalDB
	var Users models.Users
	Borcs := []models.Borcs{}

	err := c.ShouldBindJSON(&Users)
	if err != nil {
		c.JSON(200, gin.H{
			"some issue": err,
		})
		c.Abort()
	}
	db.Where("Username = ?", Users.Username).Find(&Users)
	a := Users.ID
	db.Preload("Users").Find(&Borcs, db.Where("Uid = ? ", a))
	c.JSON(200, gin.H{
		"username": Borcs,
	})
	// db.Preload("Users").Find(&Borcs, db.Where("Uid = ? ", 1))
	// c.JSON(200, gin.H{
	// 	"username": Borcs,
	// })
}
