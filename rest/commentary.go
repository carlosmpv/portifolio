package rest

import (
	"portifolio/database"
	"portifolio/forms"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ListCommentaries list every commentary api
func ListCommentaries(c *gin.Context, db *gorm.DB) {
	var commentaries []*database.Commentary

	db.Order("relevance desc").Find(&commentaries)
	c.JSON(200, commentaries)
}

// TopCommentaries get most relevant commentaries api
func TopCommentaries(c *gin.Context, db *gorm.DB) {
	var commentaries []*database.Commentary

	db.Limit(10).Order("relevance desc").Find(&commentaries)
	c.JSON(200, commentaries)
}

// CreateEditCommentary create commentary api
func CreateEditCommentary(c *gin.Context, db *gorm.DB) {
	comment := &forms.CommentaryForm{}
	err := c.ShouldBindJSON(&comment)

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	commentary := &database.Commentary{}
	commentary.Email = comment.Email
	db.Where(commentary).First(commentary)

	if commentary.ID > 0 {
		commentary.Author = comment.Author
		commentary.Content = comment.Content
		db.Save(commentary)
		c.JSON(200, gin.H{"Message": "Successfully edited"})
		return
	}

	commentary.Author = comment.Author
	commentary.Content = comment.Content
	db.Create(commentary)
	c.JSON(200, gin.H{"Message": "Successfully commented"})
}
