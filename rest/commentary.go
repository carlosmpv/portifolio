package rest

import (
	"portifolio/database"
	"portifolio/forms"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type filter struct {
	Author   string
	Content  string
	PostDate time.Time
}

func filterCommentariesAttrs(commentaries []*database.Commentary) []*filter {
	filtred := make([]*filter, len(commentaries))

	for i, co := range commentaries {
		filtred[i] = &filter{
			Author:   co.Author,
			Content:  co.Content,
			PostDate: co.CreatedAt,
		}
	}

	return filtred
}

// ListCommentaries list every commentary api
func ListCommentaries(c *gin.Context, db *gorm.DB) {
	var commentaries []*database.Commentary
	db.Order("relevance desc").Find(&commentaries)

	c.JSON(200, filterCommentariesAttrs(commentaries))
}

// TopCommentaries get most relevant commentaries api
func TopCommentaries(c *gin.Context, db *gorm.DB) {
	var commentaries []*database.Commentary
	db.Limit(10).Order("relevance desc").Find(&commentaries)

	c.JSON(200, filterCommentariesAttrs(commentaries))
}

// CreateEditCommentary create commentary api
func CreateEditCommentary(c *gin.Context, db *gorm.DB) {
	// x, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(x))

	var comment forms.CommentaryForm
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
