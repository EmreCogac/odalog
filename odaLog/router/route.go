package router

import (
	"log"
	"odaLog/odaLog/controllers"
	"odaLog/odaLog/database"
	"odaLog/odaLog/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("env not working ", err)

	}

	database.ConnectDB(&config)

}

func RouterConfig() *gin.Engine {

	r := gin.Default()
	r.GET("/getUsers", controllers.GetUsers)
	r.GET("/getBorcs", controllers.GetUserBorcs)
	r.POST("/Borcs", controllers.GetUsersFromName)
	r.GET("/", Ct)
	return r

}
func Ct(c *gin.Context) {
	c.JSON(200, gin.H{
		"a": "a",
	})
}
