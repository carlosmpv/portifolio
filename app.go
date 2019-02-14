package main

import (
	"portifolio/database"
	"portifolio/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	database := database.GetDatabase()

	server.Static("/view", "client/view")
	server.Static("/styles", "client/resources/styles")
	server.Static("/scripts", "client/resources/scripts")
	server.Static("/libs", "client/resources/bower_components")
	server.Static("/images", "client/resources/images")

	server.GET("/", func(c *gin.Context) { c.Redirect(302, "/view") })

	server.POST("/commentary/comment", func(c *gin.Context) {
		rest.CreateEditCommentary(c, database)
	})

	server.GET("/commentary/list", func(c *gin.Context) {
		rest.ListCommentaries(c, database)
	})

	server.GET("/commentary/top", func(c *gin.Context) {
		rest.TopCommentaries(c, database)
	})

	server.Run(":8080")
}
