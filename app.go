package main

import (
	"portifolio/database"
	"portifolio/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db := database.GetDatabase()
	database.MigrateAll(db)

	server.Static("/view", "client/view")
	server.Static("/styles", "client/resources/styles")
	server.Static("/scripts", "client/resources/scripts")
	server.Static("/libs", "client/resources/bower_components")
	server.Static("/images", "client/resources/images")

	server.GET("/", func(c *gin.Context) { c.Redirect(302, "/view") })
	commentary := server.Group("commentary")

	commentary.POST("/comment", func(c *gin.Context) {
		rest.CreateEditCommentary(c, db)
	})

	commentary.GET("/list", func(c *gin.Context) {
		rest.ListCommentaries(c, db)
	})

	commentary.GET("/top", func(c *gin.Context) {
		rest.TopCommentaries(c, db)
	})

	server.Run(":8080")
}
